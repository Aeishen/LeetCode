/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package binarytree

import (
	"fmt"
	"strconv"
	"strings"
)

// Codec 结构体
type Codec struct {
}

// Constructor 结构体构造函数
func Constructor() Codec {
	return Codec{}
}

// serialize 序列化
func (m *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var s string

	// 前序遍历
	var f func(*TreeNode)
	f = func(t *TreeNode) {
		if t == nil {
			s += "#,"
			return
		}
		s += strconv.Itoa(t.Val) + ","
		f(t.Left)
		f(t.Right)
	}
	f(root)
	fmt.Println(s)
	return s
}

// deserialize 反序列化
func (m *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	s := strings.Split(data, ",")
	var f func() *TreeNode
	f = func() *TreeNode {
		if s[0] == "#" {
			s = s[1:]
			return nil
		}

		val, _ := strconv.Atoi(s[0])
		t := &TreeNode{Val: val}
		s = s[1:]
		t.Left = f()
		t.Right = f()
		return t
	}

	return f()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
