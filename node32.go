package art

import "unsafe"

type Node32 struct {
	header *byte

	leaf unsafe.Pointer

	keys     *byte // [32]byte, could use AVX
	children *[32]unsafe.Pointer
}
