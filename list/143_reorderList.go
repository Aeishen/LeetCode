package list

// 重排链表

// reorderList 寻找链表中点 + 链表逆序 + 合并链表
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	re := slow.Next
	slow.Next = nil
	var p *ListNode
	for cur := re; cur != nil; {
		next := cur.Next
		cur.Next = p
		p = cur
		cur = next
	}

	for slow, fast = head, p; slow != nil || fast != nil; {
		slowNext := slow.Next
		fastNext := fast.Next

		slow.Next = fast
		fast.Next = slowNext

		slow = slowNext
		fast = fastNext
	}
}

// reorderList 线性表
func reorderList1(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	nodes := []*ListNode{}
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}
