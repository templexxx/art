package artree

import (
	"unsafe"
)

type Node2 struct {
	header *byte

	leaf unsafe.Pointer

	keys     [2]byte
	children *[2]unsafe.Pointer
}

type Node5 struct {
	header *byte

	leaf unsafe.Pointer

	keys     [5]byte
	children *[5]unsafe.Pointer
}

type Node16 struct {
	header *byte

	leaf unsafe.Pointer

	keys     *byte // [16]byte, could use SSE
	children *[16]unsafe.Pointer
}

type Node32 struct {
	header *byte

	leaf unsafe.Pointer

	keys     *byte // [32]byte, could use AVX
	children *[32]unsafe.Pointer
}

type Node64 struct {
	header *byte

	leaf unsafe.Pointer

	indexes  *byte // [256]byte, indicating char -> children[i]
	children *[64]unsafe.Pointer
}

type Node256 struct {
	header *byte

	leaf unsafe.Pointer

	children *[256]unsafe.Pointer
}

// initNode creates a Node2,
// it's the minimum Node.
func initNode() *Node2 {
	n := &Node2{
		header:   newNodeHeader(node2Type),
		leaf:     nil,
		keys:     [2]byte{},
		children: nil,
	}
	return n
}

// getNodeHeader gets header pointer from node pointer.
// header is the first field of every node type, so it will work.
func getNodeHeader(p unsafe.Pointer) *byte {
	return *(**byte)(unsafe.Pointer(p))
}
