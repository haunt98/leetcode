package main

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob_1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := root.Left
	right := root.Right

	resultRoot := root.Val

	if left != nil {
		resultRoot += rob_1(left.Left) + rob_1(left.Right)
	}

	if right != nil {
		resultRoot += rob_1(right.Left) + rob_1(right.Right)
	}

	resultChildren := rob_1(root.Left) + rob_1(root.Right)

	if resultRoot > resultChildren {
		return resultRoot
	}

	return resultChildren
}
