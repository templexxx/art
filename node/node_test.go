package node

import (
	"testing"
	"unsafe"
)

func TestGetNodeHeader(t *testing.T) {
	n2 := &Node2{
		header:   makeNodeHeader(node2Type, 0),
		leaf:     nil,
		keys:     [2]byte{},
		children: nil,
	}

	if getNodeType(getNodeHeader(unsafe.Pointer(n2))) != node2Type {
		t.Fatal("node2 type mismatch")
	}

	n5 := &Node5{
		header:   makeNodeHeader(node5Type, 0),
		leaf:     nil,
		keys:     [5]byte{},
		children: nil,
	}

	if getNodeType(getNodeHeader(unsafe.Pointer(n5))) != node5Type {
		t.Fatal("node5 type mismatch")
	}

	keys := make([]byte, 256)

	n16 := &Node16{
		header:   makeNodeHeader(node16Type, 0),
		leaf:     nil,
		keys:     &keys[:16][0],
		children: nil,
	}

	if getNodeType(getNodeHeader(unsafe.Pointer(n16))) != node16Type {
		t.Fatal("node16 type mismatch")
	}

	n32 := &Node32{
		header:   makeNodeHeader(node32Type, 0),
		leaf:     nil,
		keys:     &keys[:32][0],
		children: nil,
	}

	if getNodeType(getNodeHeader(unsafe.Pointer(n32))) != node32Type {
		t.Fatal("node32 type mismatch")
	}

	n64 := &Node64{
		header:   makeNodeHeader(node64Type, 0),
		leaf:     nil,
		indexes:  &keys[:][0],
		children: nil,
	}

	if getNodeType(getNodeHeader(unsafe.Pointer(n64))) != node64Type {
		t.Fatal("node64 type mismatch")
	}

	n256 := &Node256{
		header:   makeNodeHeader(node256Type, 0),
		leaf:     nil,
		children: nil,
	}

	if getNodeType(getNodeHeader(unsafe.Pointer(n256))) != node256Type {
		t.Fatal("node256 type mismatch")
	}
}

func TestNodeIsLeaf(t *testing.T) {
	var v uint64 = 1
	n2 := &Node2{
		header:   makeNodeHeader(node2Type, 0),
		leaf:     unsafe.Pointer(&v),
		keys:     [2]byte{},
		children: nil,
	}

	if !hasLeaf(unsafe.Pointer(n2)) {
		t.Fatal("should has leaf")
	}
	n2.leaf = nil
	if hasLeaf(unsafe.Pointer(n2)) {
		t.Fatal("shouldn't have leaf")
	}
}
