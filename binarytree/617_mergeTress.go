package binarytree

// 617. 合并二叉树

// mergeTrees 深度优先搜索
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val = t1.Val + t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

// mergeTrees1 广度优先搜索
func mergeTrees1(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	root := &TreeNode{t1.Val + t2.Val, nil, nil}
	mergeArr := []*TreeNode{root}
	t1Arr := []*TreeNode{t1}
	t2Arr := []*TreeNode{t2}

	for len(t1Arr) > 0 && len(t2Arr) > 0 {
		n := mergeArr[0]
		n1 := t1Arr[0]
		n2 := t2Arr[0]
		mergeArr = mergeArr[1:]
		t1Arr = t1Arr[1:]
		t2Arr = t2Arr[1:]

		l1 := n1.Left
		l2 := n2.Left
		r1 := n1.Right
		r2 := n2.Right
		if l1 != nil || l2 != nil {
			if l1 != nil && l2 != nil {
				left := &TreeNode{l1.Val + l2.Val, nil, nil}
				n.Left = left
				mergeArr = append(mergeArr, left)
				t1Arr = append(t1Arr, l1)
				t2Arr = append(t2Arr, l2)
			} else if l1 != nil {
				n.Left = l1
			} else {
				n.Left = l2
			}
		}
		if r1 != nil || r2 != nil {
			if r1 != nil && r2 != nil {
				right := &TreeNode{r1.Val + r2.Val, nil, nil}
				n.Right = right
				mergeArr = append(mergeArr, right)
				t1Arr = append(t1Arr, r1)
				t2Arr = append(t2Arr, r2)
			} else if r1 != nil {
				n.Right = r1
			} else {
				n.Right = r2
			}
		}
	}
	return root
}
