package nodes

import "unsafe"

type Node5 struct {
	header *byte

	leaf unsafe.Pointer

	keys     [5]byte
	children *[5]unsafe.Pointer
}
