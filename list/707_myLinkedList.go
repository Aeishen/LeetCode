package list

// MyLinkedList 我的链表
type MyLinkedList struct {
	list *ListNode
}

// Constructor Initialize your data structure here
func Constructor() MyLinkedList {
	var l MyLinkedList
	l.list = nil
	return l
}

//Get 获取链表中第 index 个节点的值。如果索引无效，则返回-1。
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || this.list == nil {
		return -1
	}

	for i, cur := 0, this.list; cur != nil; i++ {
		if i == index {
			return cur.Val
		}
		cur = cur.Next
	}
	return -1
}

//AddAtHead 在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点
func (this *MyLinkedList) AddAtHead(val int) {
	this.list = &ListNode{val, this.list}
	return
}

//AddAtTail  将值为 val 的节点追加到链表的最后一个元素
func (this *MyLinkedList) AddAtTail(val int) {
	cur := this.list
	if cur == nil {
		this.list = &ListNode{val, this.list}
		return
	}
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &ListNode{val, nil}
	return
}

//AddAtIndex 在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		this.AddAtHead(val)
		return
	}
	temp := &ListNode{0, this.list}
	cur := temp
	if cur == nil {
		return
	}
	for i := 0; cur != nil; i++ {
		if i == index {
			break
		}
		cur = cur.Next
	}
	if cur != nil {
		cur.Next = &ListNode{val, cur.Next}
	}
	this.list = temp.Next
	return
}

//DeleteAtIndex 如果索引 index 有效，则删除链表中的第 index 个节点。
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || this.list == nil {
		return
	}

	temp := &ListNode{0, this.list}
	cur := temp
	for i := 0; cur != nil; i++ {
		if i == index {
			break
		}
		cur = cur.Next
	}
	if cur != nil {
		if cur.Next != nil {
			cur.Next = cur.Next.Next
		}
	}
	this.list = temp.Next
	return
}
