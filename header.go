package artree

import (
	"encoding/binary"

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

// newNodeHeader creates a new header with specified node.
func newNodeHeader(nodeType uint32) *byte {
	p := mem.MakeAlignedBlock(16, 16)
	binary.LittleEndian.PutUint32(p[:4], nodeType)
	return &p[0]
}

// getNodeType gets node type from header.
func getNodeType(h *byte) uint8 {

	hb := mem.AtomicLoad16B(h)
	return uint8(binary.LittleEndian.Uint32(hb[:4]) & 7)
}
