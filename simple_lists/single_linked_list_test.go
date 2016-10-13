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
