package node

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
func getNodeHeader(node unsafe.Pointer) *byte {
	return *(**byte)(node)
}

// hasLeaf returns true if the node has leaf.
func hasLeaf(node unsafe.Pointer) bool {
	return getLeaf(node) != nil
}

// getLeaf gets node's leaf.
func getLeaf(node unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(*(**uint8)(unsafe.Pointer(uintptr(node) + 8)))
}

func checkPrefixPessimistic(node unsafe.Pointer, key []byte, level uint32, nonMatchingKey uint8, header *byte) {

}
