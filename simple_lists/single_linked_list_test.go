package simple_lists

import "testing"

func TestAdd(t *testing.T) {
	var node Node
	var list SingleLinkedList
	node.value = "hello"
	for _, c := range []struct {
		want string
		node Node
	}{
		{"hello", node},
	} {
		list.add(&node)
		if list.head.value != c.want || list.tail.value != c.want {
			t.Errorf("test failed, expected %s, got %s, %s", c.want, list.head.value, list.tail.value)
		}
	}
}

func TestFind(t *testing.T) {
	var node Node
	var list SingleLinkedList
	node.value = "hello"
	for _, c := range []struct {
		want string
		node Node
	}{
		{"hello", node},
	} {
		list.add(&node)
		node := list.find("hello")
		if node.value != c.want {
			t.Errorf("test failed, expected %s, got %s", c.want, node.value)
		}
	}
}
