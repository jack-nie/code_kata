package simple_lists

type SingleLinkedList struct {
	head, tail *Node
	length int
}

type Node struct {
	value string
	next *Node
}

func (list *SingleLinkedList) add(node *Node) {
	if list.head == nil{
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