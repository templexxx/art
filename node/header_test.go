package node

import "testing"

func TestNewHeader(t *testing.T) {
	h := makeNodeHeader(node2Type, 0)
	if getNodeType(h) != node2Type {
		t.Fatal("new header node type mismatch")
	}
}

func TestIsObsolete(t *testing.T) {
	h := makeNodeHeader(node2Type, 0)
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
}

func TestGetLevel(t *testing.T) {
	for i := 0; i < 128; i++ {
		h := makeNodeHeader(node2Type, uint32(i))
		lvl := getLevel(h)

		if lvl != uint32(i) {
			t.Fatal("level mismatch")
		}
	}
}
