package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := new(TreeNode)
	tree.Val = 1
	tree.Left = new(TreeNode)
	tree.Right = new(TreeNode)
	tree.Left.Val = 2
	tree.Right.Val = 3
	fmt.Println(sumNumbers(tree))
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	stack := make([]*TreeNode, 0)
	current := new(TreeNode)
	total := 0
	stack = append(stack, root)
	for len(stack) != 0 {
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if current.Right != nil {
			current.Right.Val += current.Val * 10
			stack = append(stack, current.Right)
		}
		if current.Left != nil {
			current.Left.Val += current.Val * 10
			stack = append(stack, current.Left)
		}
		if current.Left == nil && current.Right == nil {
			total += current.Val
		}
	}
	return total
}
