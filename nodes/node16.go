package nodes

import "unsafe"

type Node16 struct {
	Header *byte

	Leaf unsafe.Pointer

	Keys     *byte // [16]byte, could use SSE
	Children *[16]unsafe.Pointer
}
