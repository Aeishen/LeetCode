package list

//ListNode 链表结构
type ListNode struct {
	Val  int
	Next *ListNode
}

// 带随机指针的链表
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 多级双向链表
type Node1 struct {
	Val   int
	Prev  *Node1
	Next  *Node1
	Child *Node1
}
