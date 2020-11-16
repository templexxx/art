package nodes

import "unsafe"

type Node5 struct {
	Header *byte

	Leaf unsafe.Pointer

	Keys     [5]byte
	Children *[5]unsafe.Pointer
}
