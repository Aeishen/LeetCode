package binarytree

//二叉树剪枝

// pruneTree 后序遍历框架
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)

	if root.Left == nil && root.Right == nil {
		if root.Val == 0 {
			return nil
		}
	}
	return root
}
