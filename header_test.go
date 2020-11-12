package artree

import "testing"

func TestNewHeader(t *testing.T) {
	h := newNodeHeader(node2Type)
	if getNodeType(h) != node2Type {
		t.Fatal("new header node type mismatch")
	}
}
