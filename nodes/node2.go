package nodes

import "unsafe"

type Node2 struct {
	Header *byte

	Leaf unsafe.Pointer

	Keys     [2]byte
	Children *[2]unsafe.Pointer
}
