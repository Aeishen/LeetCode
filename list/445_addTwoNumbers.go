package list

// 两数相加 II

// addTwoNumbers 栈
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := pushStack(l1)
	s2 := pushStack(l2)
	var newHead *ListNode
	var v int
	for len(s1) > 0 || len(s2) > 0 {
		v = popStack(&s1) + popStack(&s2) + v
		newHead = &ListNode{v % 10, newHead}
		v = v / 10
	}
	if v > 0 {
		newHead = &ListNode{v, newHead}
	}
	return newHead
}

func pushStack(l *ListNode) []*ListNode {
	s := make([]*ListNode, 0)
	for cur := l; cur != nil; cur = cur.Next {
		s = append(s, cur)
	}
	return s
}
func popStack(s *[]*ListNode) int {
	n := len(*s)
	if n > 0 {
		v := (*s)[n-1].Val
		*s = (*s)[:n-1]
		return v
	}
	return 0
}
