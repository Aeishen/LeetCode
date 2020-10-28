package binarytree

//二叉树的坡度
// findTilt

var n int

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	n = 0
	traverse(root)
	return n
}

func traverse(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := traverse(root.Left)
	r := traverse(root.Right)

	n += abs(l, r)
	return root.Val + l + r
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
