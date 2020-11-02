package artree

import "unsafe"

const ()

// MakeLeafNode makes node pointer first bit to 1
// by pointer += 1.
//
// According to Go alignment guarantees(Address alignment guarantees)：
// 1. For a variable x of struct type: unsafe.Alignof(x) is the largest of all the values
// 2. If the alignment guarantee of a type T is N,
//    then the address of every value of type T must be a multiple of N at run time.
//
// And it will save 1bit memory space.
func MakeLeafNode(p unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) + 1)
}

// IsLeafNode returns the node is leaf node or not,
// returns true if it is.
func IsLeafNode(p unsafe.Pointer) bool {
	return uintptr(p)&1 == 1
}
