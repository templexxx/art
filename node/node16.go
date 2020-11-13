package node

import "unsafe"

type Node16 struct {
	header *byte

	leaf unsafe.Pointer

	keys     *byte // [16]byte, could use SSE
	children *[16]unsafe.Pointer
}
