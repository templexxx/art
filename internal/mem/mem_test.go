package mem

import (
	"testing"
	"unsafe"
)

func TestAlignedBlock(t *testing.T) {
	for i := 1; i < 33; i++ {
		b := MakeAlignedBlock(i, 16)
		if uintptr(unsafe.Pointer(&b[0]))&15 != 0 {
			t.Fatal("aligned mismatch")
		}
	}
}
