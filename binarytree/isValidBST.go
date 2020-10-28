package binarytree

import "math"

//面试题 04.05. 合法二叉搜索树
// isValidBST
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isValid(root.Left, math.MinInt32, root.Val) && isValid(root.Right, root.Val, math.MaxInt32)

}

func isValid(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return isValid(root.Left, min, root.Val) && isValid(root.Right, root.Val, max)
}
