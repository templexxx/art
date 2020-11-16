package art

import (
	"unsafe"

	"github.com/templexxx/art/nodes"

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
		// Init with Node256Type, because we assume the root will grow up quickly,
		// much faster than other inner nodes at least.
		// For avoiding keeping creating new nodes, we set the root with biggest inner node.
		root: nodes.CreateNode(nodes.Node256Type, 0),
	}
}

// Insert inserts new entry, returns nil if succeed.
//
// set insertOnly true, if you don't want update old entry.
func (t *ART) Insert(key []byte, value unsafe.Pointer, insertOnly bool) (err error) {

restart:
	needRestart := false
	var node, nextNode, parentNode unsafe.Pointer
	node = nil
	nextNode = t.root
	parentNode = nil

	var nodeKey, parentKey uint8
	var level uint32 = 0

	for {
		parentNode = node
		parentKey = nodeKey
		node = nextNode

		header := nodes.LoadNodeHeader(node)
		nexLvl := level

		var nonMatchingKey uint8

		var prefixLen uint8 = 0
		prefix := make([]byte, 12)

		switch nodes.CheckPrefix() {
		case nodes.PrefixSkippedLevel:
			goto restart
		case nodes.PrefixMismatch:

		default: // Match.

		}

	}
}

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
