package list

// 链表的中间结点

//middleNode
func middleNode(head *ListNode) *ListNode {
	pre := head
	cur := head
	for cur != nil && cur.Next != nil {
		cur = cur.Next.Next
		pre = pre.Next
	}
	return pre
}
