package artree

import "unsafe"

type Node4 struct {
	*Node
	keys     [4]byte
	children [4]unsafe.Pointer
}

type Node16 struct {
	*Node
	keys     [16]byte
	children [16]unsafe.Pointer
}

type Node32 struct {
	*Node
	keys     [32]byte
	children [32]unsafe.Pointer
}

type Node48 struct {
	*Node
	indexes  [48]byte
	children [48]unsafe.Pointer
}

type Node256 struct {
	*Node
	children [256]unsafe.Pointer
}
