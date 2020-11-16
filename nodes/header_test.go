package nodes

import (
	"bytes"
	"testing"
)

func TestMakeNodeHeader(t *testing.T) {

	for i := 0; i < 256; i++ {
		for j := node2Type; j <= node256Type; j++ {
			h := makeNodeHeader(uint8(j), uint32(i))

			hl := LoadHeader(h)

			if getNodeType(hl) != uint8(j) {
				t.Fatal("new header node type mismatch")
			}
			if getLevel(hl) != uint32(i) {
				t.Fatal("new header level mismatch")
			}
		}
	}
}

func TestIsObsolete(t *testing.T) {
	h := makeNodeHeader(node2Type, 1)
	hl := LoadHeader(h)

	if isObsolete(hl) {
		t.Fatal("node should not be obsolete")
	}

	if !setObsolete(h, hl) {
		t.Fatal("set obsolete failed")
	}

	if !isObsolete(h) {
		t.Fatal("node should be obsolete")
	}

	hl = LoadHeader(h)
	if getNodeType(hl) != node2Type {
		t.Fatal("node type mismatch")
	}
	if getLevel(hl) != 1 {
		t.Fatal("node level mismatch")
	}
	if !isObsolete(hl) {
		t.Fatal("node should be obsolete")
	}
}

func TestSetPrefix(t *testing.T) {
	h := makeNodeHeader(node2Type, 3)

	hl := LoadHeader(h)
	if !setPrefix(h, hl, []byte{1, 2}) {
		t.Fatal("set prefix should be ok")
	}
	pl := getPrefixLen(hl)
	if pl != 0 {
		t.Fatal("prefix length mismatch")
	}
	prefix := getPrefix(hl, pl)
	if prefix != nil {
		t.Fatal("prefix mismatch", prefix)
	}

	hl = LoadHeader(h)
	pl = getPrefixLen(hl)
	if pl != 2 {
		t.Fatal("prefix length mismatch")
	}
	prefix = getPrefix(hl, pl)
	if !bytes.Equal(prefix, []byte{1, 2}) {
		t.Fatal("prefix mismatch")
	}

	if getNodeType(hl) != node2Type {
		t.Fatal("node type mismatch")
	}
	if getLevel(hl) != 3 {
		t.Fatal("node level mismatch")
	}
	if isObsolete(hl) {
		t.Fatal("node should not be obsolete")
	}
}

func TestLock(t *testing.T) {

	h := makeNodeHeader(node2Type, 4)

	hl := LoadHeader(h)
	if !lock(h, hl) {
		t.Fatal("lock should be ok")
	}

	hl = LoadHeader(h)
	if !isLocked(h) {
		t.Fatal("should be locked")
	}

	hl = LoadHeader(h)
	if !unlock(h, hl) {
		t.Fatal("unlock should be ok")
	}

	hl = LoadHeader(h)
	if isLocked(h) {
		t.Fatal("should be unlocked")
	}
}
