package artree

import (
	"unsafe"
)

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
	return *(**byte)(p)
}

// hasLeaf returns true if the node has leaf.
func hasLeaf(p unsafe.Pointer) bool {
	return getLeaf(p) != nil
}

// getLeaf gets node's leaf.
func getLeaf(p unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(*(**uint8)(unsafe.Pointer(uintptr(p) + 8)))
}
