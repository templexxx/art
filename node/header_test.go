package node

import (
	"bytes"
	"testing"
)

func TestMakeNodeHeader(t *testing.T) {

	for i := 0; i < 256; i++ {
		for j := node2Type; j <= node256Type; j++ {
			h := makeNodeHeader(uint8(j), uint32(i))

			hl := load(h)

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
	hl := load(h)

	if isObsolete(hl) {
		t.Fatal("node should not be obsolete")
	}

	if !setObsolete(h, hl) {
		t.Fatal("set obsolete failed")
	}

	if !isObsolete(h) {
		t.Fatal("node should be obsolete")
	}

	hl = load(h)
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

	hl := load(h)
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

	hl = load(h)
	pl = getPrefixLen(hl)
	if pl != 2 {
		t.Fatal("prefix length mismatch")
	}
	prefix = getPrefix(hl, pl)
	if !bytes.Equal(prefix, []byte{1,2}) {
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
