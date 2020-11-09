package list

//复制带随机指针的链表
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	//复制每一个节点，使得复制后的节点都在当前节点的下一个节点
	ptr := head
	for ptr != nil {
		newNode := &Node{ptr.Val, nil, nil}
		newNode.Next = ptr.Next
		ptr.Next = newNode
		ptr = newNode.Next
	}

	//原生链表的节点的指向任意节点，使复制的节点也都指向某一任意节点
	ptr = head
	for ptr != nil {
		if ptr.Random != nil {
			ptr.Next.Random = ptr.Random.Next
		} else {
			ptr.Next.Random = nil
		}

		ptr = ptr.Next.Next
	}

	//重新连接节点，把原生节点重新连接起来，把克隆后的节点连接起来
	ptr_old_list := head
	ptr_new_list := head.Next
	head_old := head.Next
	for ptr_old_list != nil {
		ptr_old_list.Next = ptr_old_list.Next.Next
		if ptr_new_list.Next != nil {
			ptr_new_list.Next = ptr_new_list.Next.Next
		} else {
			ptr_new_list.Next = nil
		}
		ptr_old_list = ptr_old_list.Next
		ptr_new_list = ptr_new_list.Next
	}
	return head_old
}
