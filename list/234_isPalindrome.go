package list

//回文链表

//isPalindrome
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	resvers, slow, fast := head, head.Next, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		t := slow.Next
		slow.Next = resvers
		resvers = slow
		slow = t
	}

	// 奇数
	if fast == nil {
		resvers = resvers.Next
	}
	for slow != nil && resvers != nil {
		if slow.Val != resvers.Val {
			return false
		}

		slow = slow.Next
		resvers = resvers.Next
	}
	return true
}
