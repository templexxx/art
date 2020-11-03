package artree

import "unsafe"

const (
	VersionLen = 8
	ParentLen  = 8
	HeaderLen  = VersionLen + ParentLen
)

type Node struct {
	Header uint32
	Prefix [8]byte
	Parent unsafe.Pointer
}
