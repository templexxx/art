package nodes

import (
	"unsafe"
)

// CreateNode creates a new node by nodeType and its level.
func CreateNode(nodeType uint8, level uint32) unsafe.Pointer {
	switch nodeType {

	case Node2Type:
		return unsafe.Pointer(&Node2{
			Header:   makeNodeHeader(Node2Type, level),
			Leaf:     nil,
			Children: nil,
		})

	case Node5Type:
		return unsafe.Pointer(&Node5{
			Header:   makeNodeHeader(Node5Type, level),
			Leaf:     nil,
			Children: nil,
		})

	case Node16Type:
		return unsafe.Pointer(&Node16{
			Header:   makeNodeHeader(Node16Type, level),
			Leaf:     nil,
			Keys:     nil,
			Children: nil,
		})

	case Node32Type:
		return unsafe.Pointer(&Node32{
			Header:   makeNodeHeader(Node32Type, level),
			Leaf:     nil,
			Keys:     nil,
			Children: nil,
		})
	case Node64Type:
		return unsafe.Pointer(&Node64{
			Header:   makeNodeHeader(Node64Type, level),
			Leaf:     nil,
			Indexes:  nil,
			Children: nil,
		})
	default: // Node256Type.
		return unsafe.Pointer(&Node256{
			Header:   makeNodeHeader(Node256Type, level),
			Leaf:     nil,
			Children: nil,
		})
	}
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

const (
	PrefixMatch = iota
	PrefixMismatch
	PrefixSkippedLevel
)

// CheckPrefix checks prefix in a pessimistic way:
// Compare search key before proceeding to the next child.
//
// For supporting optimistic way, we have to keep the whole key in memory too,
// I prefer to save memory, although it needs more comparisons.
// (Actually, the pure pessimistic way may waste more memory, if the keys are too sparse.)
func CheckPrefix(node unsafe.Pointer, key []byte, level uint32, nonMatchingKey uint8, remainPrefix []byte) uint8 {
	return PrefixMismatch
}
