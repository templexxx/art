package artree

import "unsafe"

type Node2 struct {
	keys     [2]byte
	children *[2]unsafe.Pointer
	header   *byte
}

type Node5 struct {
	keys     [5]byte
	children *[5]unsafe.Pointer
	header   *byte
}

type Node16 struct {
	keys     [16]byte
	children *[16]unsafe.Pointer
	header   *byte
}

type Node32 struct {
	indexes  [32]byte
	children *[32]unsafe.Pointer
	header   *byte
}

type Node64 struct {
	indexes  [256]byte
	children *[64]unsafe.Pointer
	header   *byte
}

type Node256 struct {
	children *[256]unsafe.Pointer
	header   *byte
}
