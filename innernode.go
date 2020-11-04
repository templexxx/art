package artree

import "unsafe"

type Node2 struct {
	*Node
	keys     [2]byte
	children *[2]unsafe.Pointer
}

type Node5 struct {
	*Node
	keys     [5]byte
	children *[5]unsafe.Pointer
}

type Node16 struct {
	*Node
	keys     [16]byte
	children *[16]unsafe.Pointer
}

type Node32 struct {
	*Node
	indexes  [16]byte
	children *[32]unsafe.Pointer
}

type Node64 struct {
	*Node
	indexes  [64]byte
	children *[64]unsafe.Pointer
}

type Node256 struct {
	*Node
	children *[256]unsafe.Pointer
}
