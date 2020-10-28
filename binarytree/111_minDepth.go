package binarytree

import "math"

// minDepth 二叉树的最小深度
// 深度优先搜索的方法，遍历整棵树，记录最小深度。对于每一个非叶子节点，我们只需要分别计算其左右子树的最小叶子节点深度。
// 这样就将一个大问题转化为了小问题，可以递归地解决该问题。
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := minDepth(root.Left)
	r := minDepth(root.Right)
	//1.左孩子和有孩子都为空的情况，说明到达了叶子节点，直接返回1即可
	//这里两个节点为空，说明m1和m2为0，所以可以返回l + r + 1;
	//2.如果左孩子和由孩子其中一个为空，那么需要返回比较大的那个孩子的深度
	//这里其中一个节点为空，说明m1和m2有一个必然为0，所以可以返回l + r + 1;
	if root.Left == nil || root.Right == nil {
		return l + r + 1
	}

	//3.最后一种情况，也就是左右孩子都不为空，返回最小深度+1即可
	return min(l, r) + 1
}

// minDepth
// 深度优先搜索的方法，遍历整棵树，记录最小深度。对于每一个非叶子节点，我们只需要分别计算其左右子树的最小叶子节点深度。
// 这样就将一个大问题转化为了小问题，可以递归地解决该问题。
func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minD := math.MaxInt32
	if root.Left != nil {
		minD = min(minDepth1(root.Left), minD)
	}
	if root.Right != nil {
		minD = min(minDepth1(root.Right), minD)
	}
	return minD + 1
}

// minDepth
// 广度优先搜索的方法，遍历整棵树。
// 当我们找到一个叶子节点时，直接返回这个叶子节点的深度。广度优先搜索的性质保证了最先搜索到的叶子节点的深度一定最小。
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	s := []*TreeNode{root}
	c := []int{1}
	for i := 0; i < len(s); i++ {
		r := s[i]
		d := c[i]
		if r.Left == nil && r.Right == nil {
			return d
		}
		if r.Left != nil {
			s = append(s, r.Left)
			c = append(c, d+1)
		}
		if r.Right != nil {
			s = append(s, r.Right)
			c = append(c, d+1)
		}
	}
	return 0
}

// minDepth
// 广度优先搜索的方法，遍历整棵树。
// 当我们找到一个叶子节点时，直接返回这个叶子节点的深度。广度优先搜索的性质保证了最先搜索到的叶子节点的深度一定最小。
func minDepth3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	s := []*TreeNode{root}
	d := 1

	for len(s) > 0 {
		n := len(s)
		for i := 0; i < n; i++ {
			r := s[i]

			if r.Left == nil && r.Right == nil {
				return d
			}
			if r.Left != nil {
				s = append(s, r.Left)
			}
			if r.Right != nil {
				s = append(s, r.Right)
			}
		}
		s = s[n:]
		d++
	}

	return d
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
