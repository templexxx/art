package artree

import "testing"

func TestNewHeader(t *testing.T) {
	h := newHeader()
	if getNodeType(h) != node2Type {
		t.Fatal("new header node type mismatch")
	}
}
