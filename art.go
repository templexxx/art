package artree

import (
	"unsafe"

	"github.com/templexxx/cpu"
)

func init() {
	if !cpu.X86.HasCMPXCHG16B {
		panic("art need CMPXCHG16B feature, but not supported in this machine")
	}
}

// ART implements The Adaptive Radix Tree.
type ART struct {
	root unsafe.Pointer
	size uint64
}

func New() *ART {
	return &ART{
		root: unsafe.Pointer(initNode()),
	}
}

// Insert inserts new entry, returns nil if succeed.
//
// set insertOnly true, if you don't want update old entry.
//func (t *ART) Insert(key []byte, value unsafe.Pointer, insertOnly bool) (err error) {
//
//restart:
//	needRestart := false
//	var node, nextNode, parentNode unsafe.Pointer
//	node = nil
//	nextNode = t.root
//	parentNode = nil
//
//	var nodeKey, parentKey uint8
//	var level uint32 = 0
//
//	for {
//		parentNode = node
//		parentKey = nodeKey
//		node = nextNode
//
//		goto restart
//	}
//}

// Search searches key, returns value's pointer if found.
//
// Returns nil if not found.
func (t *ART) Search(key []byte) unsafe.Pointer {
	return nil
}

// Delete deletes entry,
// returns nil error if succeed, and deleted key & it's value will be returned too.
func (t *ART) Delete() (key []byte, value unsafe.Pointer, err error) {
	return nil, nil, err
}
