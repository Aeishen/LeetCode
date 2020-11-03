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
