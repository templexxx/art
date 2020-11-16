package nodes

import "unsafe"

type Node64 struct {
	Header *byte

	Leaf unsafe.Pointer

	Indexes  *byte // [256]byte, indicating char -> Children[i]
	Children *[64]unsafe.Pointer
}
