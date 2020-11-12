package artree

import (
	"testing"
	"unsafe"
)

func TestGetNodeHeader(t *testing.T) {
	n2 := &Node2{
		header:   newNodeHeader(node2Type),
		leaf:     nil,
		keys:     [2]byte{},
		children: nil,
	}

	if getNodeType(getNodeHeader(unsafe.Pointer(n2))) != node2Type {
		t.Fatal("node2 type mismatch")
	}

	n5 := &Node5{
		header:   newNodeHeader(node5Type),
		leaf:     nil,
		keys:     [5]byte{},
		children: nil,
	}

	if getNodeType(getNodeHeader(unsafe.Pointer(n5))) != node5Type {
		t.Fatal("node5 type mismatch")
	}

	keys := make([]byte, 256)

	n16 := &Node16{
		header:   newNodeHeader(node16Type),
		leaf:     nil,
		keys:     &keys[:16][0],
		children: nil,
	}

	if getNodeType(getNodeHeader(unsafe.Pointer(n16))) != node16Type {
		t.Fatal("node16 type mismatch")
	}

	n32 := &Node32{
		header:   newNodeHeader(node32Type),
		leaf:     nil,
		keys:     &keys[:32][0],
		children: nil,
	}

	if getNodeType(getNodeHeader(unsafe.Pointer(n32))) != node32Type {
		t.Fatal("node32 type mismatch")
	}

	n64 := &Node64{
		header:   newNodeHeader(node64Type),
		leaf:     nil,
		indexes:  &keys[:][0],
		children: nil,
	}

	if getNodeType(getNodeHeader(unsafe.Pointer(n64))) != node64Type {
		t.Fatal("node64 type mismatch")
	}

	n256 := &Node256{
		header:   newNodeHeader(node256Type),
		leaf:     nil,
		children: nil,
	}

	if getNodeType(getNodeHeader(unsafe.Pointer(n256))) != node256Type {
		t.Fatal("node256 type mismatch")
	}
}
