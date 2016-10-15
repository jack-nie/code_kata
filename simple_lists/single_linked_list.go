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
	l := list.head
	for l != nil && l.value != value {
		l = l.next
	}
	return *l
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

func (list *SingleLinkedList) isLast(node *Node) bool {
	if node == nil {
		return true
	}
	return false
}

func (list *SingleLinkedList) delete(node Node) {
	prev := list.findPrevious(node.value)
	head := list.head
	var nilNode *Node
	for !list.isLast(head) && head != prev {
		head = head.next
	}

	if head.next != nil {
		head.next = head.next.next
	} else {
		head.next = nilNode
	}
}

func (list *SingleLinkedList) findPrevious(value string) *Node {
	head := list.head
	for head.next != nil && head.next.value != value {
		head = head.next
	}

	if head.next == nil {
		return head.next
	}

	return head
}
