package list

// 奇偶链表

// oddEvenList 双指针
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	f, s := head, head.Next

	for cur := s; cur != nil && cur.Next != nil; {
		f.Next = f.Next.Next
		cur.Next = cur.Next.Next
		f = f.Next
		cur = cur.Next
	}
	f.Next = s

	return head
}
