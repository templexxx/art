package artree

import (
	"testing"
	"unsafe"
)

func TestIsLeafNode(t *testing.T) {

	for i := 0; i < 1024; i++ {
		n := new(Node)
		if IsLeafNode(unsafe.Pointer(n)) {
			t.Fatal("should not be leaf node")
		}

		ln := MakeLeafNode(unsafe.Pointer(n))
		if !IsLeafNode(ln) {
			t.Fatal("should be leaf node")
		}
	}
}
