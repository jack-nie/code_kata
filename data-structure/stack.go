package dataStructure

type Stack struct {
	nodes []*Node
	count int
}

func (s *Stack) Push(n *Node) {
	if s.count >= len(s.nodes) {
		nodes := make([]*Node, len(s.nodes)*2)
		copy(nodes, s.nodes)
		s.nodes = nodes
	}
	s.nodes[s.count] = n
	s.count++
}

func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	node := s.nodes[s.count-1]
	s.count--
	return node
}
