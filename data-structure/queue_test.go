package dataStructure

import "testing"

func testPush(t *testing.T) {
	q := &Queue{nodes: make([]*Node, 3)}
	q.Push(&Node{1})
	if q.nodes[0].Value != 1 {
		t.Fatal("expectd 1, but got %d", q.nodes[0].Value)
	}
}

func testPop(t *testing.T) {
	q := &Queue{nodes: make([]*Node, 3)}
	q.Push(&Node{1})
	v := q.Pop().Value
	if v != 1 {
		t.Fatal("expectd 1, but got %d", v)
	}
}
