package artree

import "unsafe"

const (
	VersionLen = 8
	ParentLen  = 8
	HeaderLen  = VersionLen + ParentLen
)

// Node is the basic structure of art.
type Node struct {
	// Header includes basic information of a node:
	// Struct:
	// +-----------+-------------+---------+------------+----------+----------+
	// | offset(8) | prefix_len(4) |  type(3) | old(1) | locked(1) | expand(1) | vexpand(8)
	// +-----------+-------------+---------+------------+----------+----------+
	Header uint32
	Prefix [8]byte
	Parent unsafe.Pointer
}
