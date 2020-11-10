package artree

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
