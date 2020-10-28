package binarytree

// 修剪二叉搜索树

// trimBST 深度优先
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	// 根节点小于最小，变更根节点
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}

	// 根节点大于最大，变更根节点
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}

	// 根节点大于等于最小，小于等于最大，无需变换根节点
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}
