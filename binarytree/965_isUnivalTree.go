package binarytree

//单值二叉树
// isUnivalTree 递归
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	v := root.Val
	if root.Left != nil && root.Left.Val != v {
		return false
	}
	if root.Right != nil && root.Right.Val != v {
		return false
	}

	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}
