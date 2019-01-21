package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(root, 1)
}

func max(root *TreeNode, n int) int {
	if root.Left == nil && root.Right == nil {
		return n
	}

	left := n
	right := n

	if root.Left != nil {
		left = max(root.Left, n+1)
	}

	if root.Right != nil {
		right = max(root.Right, n+1)
	}

	if left > right {
		return left
	}
	return right
}
