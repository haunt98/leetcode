package main

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST_1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	leftResult := searchBST_1(root.Left, val)
	if leftResult != nil {
		return leftResult
	}

	return searchBST_1(root.Right, val)
}
