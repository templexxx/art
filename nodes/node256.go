package nodes

import "unsafe"

type Node256 struct {
	Header *byte

	Leaf unsafe.Pointer

	Children *[256]unsafe.Pointer
}
