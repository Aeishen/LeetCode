package list

// 旋转链表

// rotateRight 形成环形链表再断开
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	p := head
	l := 1
	for ; p.Next != nil; p = p.Next {
		l++
	}
	k = k % l
	p.Next = head

	for i := 0; i < l-k; i++ {
		p = p.Next
	}

	head, p.Next = p.Next, nil
	return head
}
