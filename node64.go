package artree

import "unsafe"

type Node64 struct {
	header *byte

	leaf unsafe.Pointer

	indexes  *byte // [256]byte, indicating char -> children[i]
	children *[64]unsafe.Pointer
}
