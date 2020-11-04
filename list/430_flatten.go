package list

// 扁平化多级双向链表

// flatten
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
