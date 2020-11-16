package nodes

import "unsafe"

type Node32 struct {
	Header *byte

	Leaf unsafe.Pointer

	Keys     *byte // [32]byte, could use AVX
	Children *[32]unsafe.Pointer
}
