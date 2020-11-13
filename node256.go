package art

import "unsafe"

type Node256 struct {
	header *byte

	leaf unsafe.Pointer

	children *[256]unsafe.Pointer
}
