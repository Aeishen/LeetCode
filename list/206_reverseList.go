package list

// 反转链表

//reverseList
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	pre := head
	cur := head.Next
	pre.Next = nil

	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

//reverseList
func reverseList2(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

//reverseList
func reverseList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	q := reverseList3(head.Next)
	head.Next.Next = head.Next
	head.Next = nil
	return q
}
