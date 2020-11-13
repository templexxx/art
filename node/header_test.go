package node

import "testing"

func TestNewHeader(t *testing.T) {
	h := newNodeHeader(node2Type)
	if getNodeType(h) != node2Type {
		t.Fatal("new header node type mismatch")
	}
}

func TestIsObsolete(t *testing.T) {
	h := newNodeHeader(node2Type)
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
