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
	return *(**byte)(unsafe.Pointer(p))
}
