package list

// 环形链表
// hasCycle 快慢指针
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	s := head
	f := head.Next

	for s != f {
		if f == nil || f.Next == nil {
			return false
		}
		s = s.Next
		f = f.Next.Next
	}

	return true
}

// hasCycle 快慢指针
func hasCycle1(head *ListNode) bool {

	s := head
	f := head

	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if f == s {
			return true
		}
	}

	return false
}

// hasCycle 哈希表
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	// todo
	return false
}

// hasCycle 先反转在比较
func hasCycle3(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	// todo
	return false
}

// hasCycle 递归逐个删除
func hasCycle4(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	// todo
	return false
}
