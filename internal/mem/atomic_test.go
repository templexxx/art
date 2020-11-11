package mem

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"runtime"
	"testing"
	"unsafe"
)

var (
	magic128 = make([]byte, 16)
)

func init() {
	binary.LittleEndian.PutUint64(magic128[:8], 0xdeddeadbeefbeef)
	binary.LittleEndian.PutUint64(magic128[8:], 0xdeddeadbeefbeef)
}

func TestAtomicLoad16B(t *testing.T) {
	var x struct {
		before []uint8
		i      []byte
		after  []uint8
	}
	x.before = magic128
	x.after = magic128
	x.i = MakeAlignedBlock(16, 16)

	for delta := uint64(1); delta+delta > delta; delta += delta {
		k := AtomicLoad16B(&x.i[0])
		if !bytes.Equal(k[:], x.i) {
			t.Fatalf("delta=%d i=%d k=%d", delta, x.i, k)
		}

		xi0 := binary.LittleEndian.Uint64(x.i[:8])
		xi0 += delta
		xi1 := binary.LittleEndian.Uint64(x.i[8:])
		xi1 -= delta

		binary.LittleEndian.PutUint64(x.i[:8], delta+2)
		binary.LittleEndian.PutUint64(x.i[8:], ^(delta + 2))
	}
	if !bytes.Equal(x.before, magic128) || !bytes.Equal(x.after, magic128) {
		t.Fatal("wrong magic")
	}
}

func TestAtomicStore16B(t *testing.T) {
	var x struct {
		before []uint8
		i      []byte
		after  []uint8
	}
	x.before = magic128
	x.after = magic128
	x.i = MakeAlignedBlock(16, 16)

	var v [16]byte
	for delta := uint64(1); delta+delta > delta; delta += delta {
		AtomicStore16B(&x.i[0], v)
		if !bytes.Equal(v[:], x.i) {
			t.Fatalf("delta=%d i=%d", delta, x.i)
		}

		xi0 := binary.LittleEndian.Uint64(v[:8])
		xi0 += delta
		xi1 := binary.LittleEndian.Uint64(v[8:])
		xi1 -= delta

		binary.LittleEndian.PutUint64(v[:8], delta+2)
		binary.LittleEndian.PutUint64(v[8:], ^(delta + 2))
	}
	if !bytes.Equal(x.before, magic128) || !bytes.Equal(x.after, magic128) {
		t.Fatal("wrong magic")
	}
}

func TestAtomicCAS16B(t *testing.T) {

	var x struct {
		before []uint8
		i      []byte
		after  []uint8
	}
	x.before = magic128
	x.after = magic128

	old := make([]byte, 16)
	newV := make([]byte, 16)
	nnv := make([]byte, 16)
	x.i = MakeAlignedBlock(16, 16)

	for val := uint64(1); val+val > val; val += val {

		binary.LittleEndian.PutUint64(old[:8], val)
		binary.LittleEndian.PutUint64(old[8:], ^val)

		binary.LittleEndian.PutUint64(newV[:8], val+1)
		binary.LittleEndian.PutUint64(newV[8:], ^(val + 1))

		copy(x.i, old)

		if !AtomicCAS16B(&x.i[0], &old[0], &newV[0]) {
			t.Fatal("should have swapped")
		}
		if !bytes.Equal(x.i, newV) {
			t.Fatalf("wrong x.i after swap: x.i=%#x exp=%#x", x.i, newV)
		}

		copy(x.i, newV)

		binary.LittleEndian.PutUint64(nnv[:8], val+2)
		binary.LittleEndian.PutUint64(nnv[8:], ^(val + 2))

		if AtomicCAS16B(&x.i[0], &old[0], &nnv[0]) {
			t.Fatal("should not have swapped")
		}
		if !bytes.Equal(x.i, newV) {
			t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, newV)
		}
	}
	if !bytes.Equal(x.before, magic128) || !bytes.Equal(x.after, magic128) {
		t.Fatal("wrong magic")
	}
}

// Tests of correct behavior, with contention.
// (Is the function atomic?)
//
// For each function, we write a "hammer" function that repeatedly
// uses the atomic operation to add 1 to a value. After running
// multiple hammers in parallel, check that we end with the correct
// total.
// Swap can't add 1, so it uses a different scheme.
// The functions repeatedly generate a pseudo-random number such that
// low bits are equal to high bits, swap, check that the old value
// has low and high bits equal.

var hammer128 = map[string]func(*byte, int){
	"AtomicCAS16B": hammerAtomicCAS16B,
}

func hammerAtomicCAS16B(addr *byte, count int) {
	for i := 0; i < count; i++ {
		for {
			v := AtomicLoad16B(addr)
			v0 := binary.LittleEndian.Uint64(v[:8])
			v1 := binary.LittleEndian.Uint64(v[8:])
			var vv [16]byte
			binary.LittleEndian.PutUint64(vv[:8], v0+1)
			binary.LittleEndian.PutUint64(vv[8:], v1-1)
			if AtomicCAS16B(addr, &v[0], &vv[0]) {
				break
			}
		}
	}
}

func TestHammer128(t *testing.T) {
	const p = 4
	n := 10000
	if testing.Short() {
		n = 1000
	}
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(p))

	for name, testf := range hammer128 {
		c := make(chan int)

		val := MakeAlignedBlock(16, 16)
		for i := 0; i < p; i++ {
			go func() {
				defer func() {
					if err := recover(); err != nil {
						t.Error(err.(string))
					}
					c <- 1
				}()
				testf(&val[0], n)
			}()
		}
		for i := 0; i < p; i++ {
			<-c
		}
		var exp [16]byte
		binary.LittleEndian.PutUint64(exp[:8], uint64(n)*p)
		binary.LittleEndian.PutUint64(exp[8:], -uint64(n)*p)
		if !bytes.Equal(val, exp[:]) {
			t.Fatalf("%s: val=%v want %v", name, val, exp)
		}
	}
}

func hammerStoreLoadUint128(t *testing.T, paddr unsafe.Pointer) {
	addr := (*byte)(paddr)
	v := AtomicLoad16B(addr)
	v0 := binary.LittleEndian.Uint64(v[:8])
	v1 := binary.LittleEndian.Uint64(v[8:])

	if v0 != v1 {
		t.Fatalf("Uint128: %#x != %#x", v0, v1)
	}
	var newV [16]byte
	binary.LittleEndian.PutUint64(newV[:8], v0+1)
	binary.LittleEndian.PutUint64(newV[8:], v1+1)

	AtomicStore16B(addr, newV)
}

func TestHammerStoreLoad(t *testing.T) {
	var tests []func(*testing.T, unsafe.Pointer)
	tests = append(tests, hammerStoreLoadUint128)
	n := int(1e6)
	if testing.Short() {
		n = int(1e4)
	}
	const procs = 8
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(procs))
	for _, tt := range tests {
		c := make(chan int)
		var val uint64
		for p := 0; p < procs; p++ {
			go func() {
				for i := 0; i < n; i++ {
					tt(t, unsafe.Pointer(&val))
				}
				c <- 1
			}()
		}
		for p := 0; p < procs; p++ {
			<-c
		}
	}
}

func TestNilDeref(t *testing.T) {
	funcs := [...]func(){
		func() {
			var a, b [16]byte
			AtomicCAS16B(nil, &(a[0]), &(b[0]))
		},
		func() { AtomicLoad16B(nil) },
		func() { AtomicStore16B(nil, [16]byte{0}) },
	}
	for _, f := range funcs {
		func() {
			defer func() {
				runtime.GC()
				recover()
			}()
			f()
		}()
	}
}

func BenchmarkCAS16B(b *testing.B) {

	dst := MakeAlignedBlock(16, 16)
	o := make([]byte, 16)
	n := make([]byte, 16)
	rand.Read(o)
	rand.Read(n)

	b.SetBytes(16)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = AtomicCAS16B(&dst[0], &o[0], &n[0])
	}
}
