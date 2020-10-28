package list

//合并两个排序的链表

// mergeTwoLists
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	h := new(ListNode)
	p := h

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			p.Next = l2
			l2 = l2.Next
		} else {
			p.Next = l1
			l1 = l1.Next
		}
		p = p.Next
	}

	if l2 != nil {
		p.Next = l2
	}
	if l1 != nil {
		p.Next = l1
	}
	return h.Next
}

// mergeTwoLists 递归
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil && l2 != nil {
		return l2
	}
	if l1 != nil && l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists1(l1.Next, l2)
	} else {
		l1, l2 = l2, l1
		l1.Next = mergeTwoLists1(l1.Next, l2)
	}

	return l1
}
