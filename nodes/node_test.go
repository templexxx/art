package nodes

import (
	"testing"
	"unsafe"
)

func TestGetNodeHeader(t *testing.T) {
	n2 := &Node2{
		Header:   makeNodeHeader(Node2Type, 0),
		Leaf:     nil,
		Keys:     [2]byte{},
		Children: nil,
	}

	if getNodeType(GetHeader(unsafe.Pointer(n2))) != Node2Type {
		t.Fatal("node2 type mismatch")
	}

	n5 := &Node5{
		Header:   makeNodeHeader(Node5Type, 0),
		Leaf:     nil,
		Keys:     [5]byte{},
		Children: nil,
	}

	if getNodeType(GetHeader(unsafe.Pointer(n5))) != Node5Type {
		t.Fatal("node5 type mismatch")
	}

	keys := make([]byte, 256)

	n16 := &Node16{
		Header:   makeNodeHeader(Node16Type, 0),
		Leaf:     nil,
		Keys:     &keys[:16][0],
		Children: nil,
	}

	if getNodeType(GetHeader(unsafe.Pointer(n16))) != Node16Type {
		t.Fatal("node16 type mismatch")
	}

	n32 := &Node32{
		Header:   makeNodeHeader(Node32Type, 0),
		Leaf:     nil,
		Keys:     &keys[:32][0],
		Children: nil,
	}

	if getNodeType(GetHeader(unsafe.Pointer(n32))) != Node32Type {
		t.Fatal("node32 type mismatch")
	}

	n64 := &Node64{
		Header:   makeNodeHeader(Node64Type, 0),
		Leaf:     nil,
		Indexes:  &keys[:][0],
		Children: nil,
	}

	if getNodeType(GetHeader(unsafe.Pointer(n64))) != Node64Type {
		t.Fatal("node64 type mismatch")
	}

	n256 := &Node256{
		Header:   makeNodeHeader(Node256Type, 0),
		Leaf:     nil,
		Children: nil,
	}

	if getNodeType(GetHeader(unsafe.Pointer(n256))) != Node256Type {
		t.Fatal("node256 type mismatch")
	}
}

func TestNodeIsLeaf(t *testing.T) {
	var v uint64 = 1
	n2 := &Node2{
		Header:   makeNodeHeader(Node2Type, 0),
		Leaf:     unsafe.Pointer(&v),
		Keys:     [2]byte{},
		Children: nil,
	}

	if !hasLeaf(unsafe.Pointer(n2)) {
		t.Fatal("should has leaf")
	}
	n2.Leaf = nil
	if hasLeaf(unsafe.Pointer(n2)) {
		t.Fatal("shouldn't have leaf")
	}
}
