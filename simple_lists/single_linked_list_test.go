package simple_lists

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	var list SingleLinkedList
	for _, c := range []struct {
		want  string
		value string
	}{
		{"hello", "hello"},
	} {
		list.add(c.value)
		if list.head.value != c.want || list.tail.value != c.want {
			t.Errorf("test failed, expected %s, got %s, %s", c.want, list.head.value, list.tail.value)
		}
	}
}

func TestFind(t *testing.T) {
	var list SingleLinkedList
	for _, c := range []struct {
		want  string
		value string
	}{
		{"hello", "hello"},
		{"world", "world"},
		{"!", "!"},
	} {
		list.add(c.value)
		node := list.find(c.value)
		if node.value != c.want {
			t.Errorf("test failed, expected %s, got %s", c.want, node.value)
		}
	}
}

func TestValues(t *testing.T) {
	var list SingleLinkedList
	for _, c := range []struct {
		want []string
	}{
		{[]string{"hello", "world"}},
	} {
		list.add("hello")
		list.add("world")
		got := list.values()
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("test failed, expected %s, got %s", c.want, got)
		}
	}
}

func TestIsLast(t *testing.T) {
	var list SingleLinkedList
	list.add("world")
	list.add("hello")
	for _, c := range []struct {
		want bool
		head *Node
	}{
		{false, list.head},
		{false, list.head.next},
		{true, list.head.next.next},
	} {
		got := list.isLast(c.head)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("test failed, expected %s, got %s", c.want, got)
		}
	}
}

func TestFindPrevious(t *testing.T) {
	var list SingleLinkedList
	list.add("world")
	list.add("hello")
	list.add("google")
	for _, c := range []struct {
		want  string
		value string
	}{
		{"world", "hello"},
		{"hello", "google"},
	} {
		got := list.findPrevious(c.value)

		if got != nil && got.value != c.want {
			t.Errorf("test failed, expected %s, got %s", c.want, got.value)
		}
	}
}
