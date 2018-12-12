package main

import (
	"fmt"
	"strconv"
)

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
	preOrder(tree)
	inOrder(tree)
	postOrder(tree)
}

func preOrder(root *TreeNode) {
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) != 0 {
			root = stack[len(stack)-1]
			fmt.Println(strconv.Itoa(root.Val))
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
}

func inOrder(root *TreeNode) {
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			fmt.Println(root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) != 0 {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
}

func postOrder(root *TreeNode) {
	stack := make([]*TreeNode, 0)
	var pre *TreeNode
	var cur *TreeNode
	stack = append(stack, root)
	for len(stack) != 0 {
		cur = stack[len(stack)-1]
		if (cur.Left == nil && cur.Right == nil) || (pre != nil && (cur.Left == pre || cur.Right == pre)) {
			fmt.Println(cur.Val)
			stack = stack[:len(stack)-1]
			pre = cur
		} else {
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
		}
	}
}
