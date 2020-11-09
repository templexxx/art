// Copyright (C) 2012 by Nick Craig-Wood http://www.craig-wood.com/nick/

package artree

import "unsafe"

// alignment returns alignment of the block in memory
// with reference to alignSize.
//
// Can't check alignment of a zero sized block as &block[0] is invalid.
func alignment(block []byte, alignSize int) int {
	return int(uintptr(unsafe.Pointer(&block[0])) & uintptr(alignSize-1))
}

// alignedBlock returns []byte of size BlockSize aligned to a multiple
// of alignSize in memory (must be power of two).
func alignedBlock(blockSize, alignSize int) []byte {
	block := make([]byte, blockSize+alignSize)
	if alignSize == 0 {
		return block
	}
	a := alignment(block, alignSize)
	offset := 0
	if a != 0 {
		offset = alignSize - a
	}
	block = block[offset : offset+blockSize]
	// Can't check alignment of a zero sized block.
	if blockSize != 0 {
		a = alignment(block, alignSize)
		if a != 0 {
			panic("failed to align block")
		}
	}
	return block
}
