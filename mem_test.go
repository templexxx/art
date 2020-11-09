package artree

import (
	"bytes"
	"encoding/binary"
	"testing"
)

var (
	magic128 = []uint8{1, 16, 12, 17, 17, 12, 16, 1, 1, 16, 12, 17, 16, 1, 17, 12}
)

func TestCompareAndSwap16Bytes(t *testing.T) {

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
	x.i = alignedBlock(16, 16)

	for val := uint64(1); val+val > val; val += val {

		binary.LittleEndian.PutUint64(old[:8], val)
		binary.LittleEndian.PutUint64(old[8:], val)

		binary.LittleEndian.PutUint64(newV[:8], val+1)
		binary.LittleEndian.PutUint64(newV[8:], val+1)

		copy(x.i, old)

		if !cas16b(&x.i[0], &old[0], &newV[0]) {
			t.Fatalf("should have swapped %#x %#x", val, val+1)
		}
		if !bytes.Equal(x.i, newV) {
			t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
		}

		copy(x.i, newV)

		binary.LittleEndian.PutUint64(nnv[:8], val+2)
		binary.LittleEndian.PutUint64(nnv[8:], val+2)

		if cas16b(&x.i[0], &old[0], &nnv[0]) {
			t.Fatalf("should not have swapped %#x %#x", val, val+2)
		}
		if !bytes.Equal(x.i, newV) {
			t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
		}
	}
	if !bytes.Equal(x.before, magic128) || !bytes.Equal(x.after, magic128) {
		t.Fatal("wrong magic")
	}
}
