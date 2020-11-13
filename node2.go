package art

import "unsafe"

type Node2 struct {
	header *byte

	leaf unsafe.Pointer

	keys     [2]byte
	children *[2]unsafe.Pointer
}
