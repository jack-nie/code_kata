package simple_lists

type SingleLinkedList struct {
	head, tail *Node
	length     int
}

type Node struct {
	value string
	next  *Node
}

func (list *SingleLinkedList) add(value string) {
	var node Node
	node.value = value
	if list.head == nil {
		list.head = &node
		list.tail = &node
	} else if list.head == list.tail {
		list.head.next = &node
		list.tail = &node
	} else {
		list.tail.next = &node
		list.tail = &node
	}
	list.length++
}

func (list *SingleLinkedList) find(value string) Node {
	var node Node
	l := list.head
	for {
		if l == nil {
			return node
		} else if l.value == value {
			return *l
		}
		l = l.next
	}
	return node
}

func (list *SingleLinkedList) values() []string {
	var values []string
	l := list.head
	for {
		if l == nil {
			break
		} else {
			values = append(values, l.value)
		}
		l = l.next
	}
	return values
}
