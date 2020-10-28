package list

// 删除链表中的节点
func deleteNode(node *ListNode) {
	*node = *node.Next
}
