package node

import (
	"encoding/binary"
	"unsafe"

	"github.com/templexxx/art/internal/mem"
)

// Header includes basic information of a node:
// 1. Helping to
// Struct(from lowest bit to highest bit):
// +---------+-------------+-----------+-----------+---------------+------------+
// | type(3) | obsolete(1) | locked(1) | level(23) | prefix_len(4) | prefix(96) |
// +---------+-------------+-----------+-----------+---------------+------------+
// 0                                                                           128
//
// type: node type
// obsolete: node is obsolete or not, 1 means obsolete
// locked: node is locked or not, 1 means locked
// level: node height, including prefix
// prefix_len: node prefix length
// prefix: node prefix

// TODO add header pool

const headerLen = 16

const (
	nodeTypeReserved = iota
	node2Type
	node5Type
	node16Type
	node32Type
	node64Type
	node256Type
)

// makeNodeHeader makes a new header with specified node type and its level.
// level won't change in node's whole life.
func makeNodeHeader(nodeType, level uint32) *byte {
	p := mem.MakeAlignedBlock(16, 16)
	n := nodeType | level<<5
	binary.LittleEndian.PutUint32(p[:4], n)
	return &p[0]
}

// load loads header to a new address, and return the address.
func load(p *byte) *byte {
	hb := mem.AtomicLoad16B(p)
	return &hb[0]
}

// getNodeType gets node type from header.
// h must be returned by func load.
func getNodeType(h *byte) uint8 {

	return *h & 7
}

// isObsolete returns node is obsolete or not.
// h must be returned by func load.
func isObsolete(h *byte) bool {

	if (*h>>3)&1 == 1 {
		return true
	}
	return false
}

// setObsolete sets node obsolete.
// h is the origin header address.
// old is the loaded header address.
func setObsolete(h, old *byte) bool {

	oldb := (*[16]byte)(unsafe.Pointer(old))
	newb := *oldb
	newb[0] = newb[0] | (1 << 3)
	return mem.AtomicCAS16B(h, old, &newb[0])
}

// getLevel gets node level.
// h must be returned by func load.
func getLevel(h *byte) uint32 {

	hb := (*[16]byte)(unsafe.Pointer(h))
	return binary.LittleEndian.Uint32((*hb)[:4]) >> 5 & (1<<23 - 1)
}

// getPrefix gets node prefix and its length.
// h must be returned by func load.
func getPrefix(h *byte) (prefixLen uint8, prefix *byte) {

}
