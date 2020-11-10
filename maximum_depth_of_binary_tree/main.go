package main

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth_1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth_1(root.Left)
	rightdepth := maxDepth_1(root.Right)

	if leftDepth >= rightdepth {
		return leftDepth + 1
	}

	return rightdepth + 1
}
