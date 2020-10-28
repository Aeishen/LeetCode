package list

// 移除链表元素

// removeElements
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{0, head}

	for pre, cur := dummy, dummy.Next; cur != nil; cur = pre.Next {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
	}
	return dummy.Next
}
