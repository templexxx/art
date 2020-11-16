package nodes

import (
	"unsafe"
)

// initNode creates a Node2,
// it's the minimum Node.
func initNode(level uint32) *Node2 {
	n := &Node2{
		header:   makeNodeHeader(node2Type, level),
		leaf:     nil,
		keys:     [2]byte{},
		children: nil,
	}
	return n
}

// GetHeader gets header pointer from node pointer.
// header is the first field of every node type, so it will work.
func GetHeader(node unsafe.Pointer) *byte {
	return *(**byte)(node)
}

// LoadNodeHeader loads node header to a new address, and return the address.
func LoadNodeHeader(node unsafe.Pointer) *byte {
	return LoadHeader(GetHeader(node))
}

// hasLeaf returns true if the node has leaf.
func hasLeaf(node unsafe.Pointer) bool {
	return getLeaf(node) != nil
}

// getLeaf gets node's leaf.
func getLeaf(node unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(*(**uint8)(unsafe.Pointer(uintptr(node) + 8)))
}

// checkPrefix checks prefix in a pessimistic way:
// Compare search key before proceeding to the next child.
//
// For supporting optimistic way, we have to keep the whole key in memory too,
// I prefer to save memory, although it needs more comparisons.
func checkPrefix(node unsafe.Pointer, key []byte, level uint32, nonMatchingKey uint8, header *byte) {

}
