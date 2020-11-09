package list

// 扁平化多级双向链表

// flatten 递归
func flatten(root *Node1) *Node1 {
	if root == nil || (root.Next == nil && root.Child == nil) {
		return root
	}

	cur := root
	for cur != nil {
		next := cur.Next
		reCur, isChange := dealChild(cur)
		if isChange {
			cur = reCur
			for cur.Next != nil {
				cur = cur.Next
			}
		}
		cur.Next = next
		if next != nil {
			next.Prev = cur
		}
		cur = cur.Next
	}
	return root
}

func dealChild(root *Node1) (*Node1, bool) {
	if root.Child == nil {
		return root, false
	}
	root.Next = flatten(root.Child)
	root.Child = nil
	root.Next.Prev = root
	return root, true
}

// 递归 扁平化的操作也就是对二叉树进行先序遍历（深度优先搜索）
func flatten1(root *Node1) *Node1 {
	if root == nil || (root.Next == nil && root.Child == nil) {
		return root
	}
	pseudoHead := &Node1{0, nil, root, nil}
	flattenDFS(pseudoHead, root)
	pseudoHead.Next.Prev = nil
	return pseudoHead.Next
}

func flattenDFS(prev, cur *Node1) *Node1 {
	if cur == nil {
		return prev
	}
	cur.Prev = prev
	prev.Next = cur

	tempNext := cur.Next

	tail := flattenDFS(cur, cur.Child)
	cur.Child = nil

	return flattenDFS(tail, tempNext)
}

// 迭代的深度优先搜索
func flatten2(root *Node1) *Node1 {
	if root == nil || (root.Next == nil && root.Child == nil) {
		return root
	}

	pseudoHead := &Node1{0, nil, root, nil}
	s := []*Node1{root}
	prev := pseudoHead
	for len(s) > 0 {
		cur := s[len(s)-1]
		s = s[:len(s)-1]
		prev.Next = cur
		cur.Prev = prev

		if cur.Next != nil {
			s = append(s, cur.Next)
		}

		if cur.Child != nil {
			s = append(s, cur.Child)
			cur.Child = nil
		}
		prev = cur
	}
	pseudoHead.Next.Prev = nil
	return pseudoHead.Next
}
