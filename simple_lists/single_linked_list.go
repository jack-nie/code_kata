package simple_lists

type SingleLinkedList struct {
	head, tail *Node
	length     int
}

type Node struct {
	value string
	next  *Node
}

func (list *SingleLinkedList) add(node *Node) {
	if list.head == nil {
		list.head = node
		list.tail = node
	} else if list.head == list.tail {
		list.head.next = node
		list.tail = node
	} else {
		list.tail.next = node
		list.tail = node
	}
	list.length++
}

func (list *SingleLinkedList) find(value string) Node {
	var node Node
	l := list.head
	for {
		if l.value == value {
			return *l
		} else {
			l = l.next
		}
	}
	return node
}
