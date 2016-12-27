package dataStructure

import "testing"

func testPush(t *testing.T) {
	s := &Stack{nodes: make([]*Node, 3)}
	s.Push(&Node{1})

	if s.nodes[0].Value != 1 {
		t.Fatal("expected 1, but got %d", s.nodes[0].Value)
	}
}

func testPop(t *testing.T) {
	s := &Stack{nodes: make([]*Node, 3)}
	s.Push(&Node{1})
	v := s.Pop().Value

	if v != 1 {
		t.Fatal("expected 1, but got %d", v)
	}
}
