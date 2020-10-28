package binarytree

import (
	"math"
	"strconv"
)

var s [][]string

func printTree(root *TreeNode) [][]string {
	if root == nil {
		return [][]string{}
	}

	d := getDepth(root)
	c := getCount(d)
	s = make([][]string, d)
	for i := 0; i < len(s); i++ {
		s[i] = make([]string, c)
		for j := 0; j < c; j++ {
			s[i][j] = ""
		}
	}
	getResult(root, 0, d, c, 0)
	return s
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + int(math.Max(float64(getDepth(root.Left)), float64(getDepth(root.Right))))
}

func getCount(i int) int {
	return int(math.Pow(2, float64(i))) - 1
}

func getResult(root *TreeNode, start, d, en, index int) {
	if index == d {
		return
	}
	mid := len(s[index][start:en]) / 2
	s[index][start:en][mid] = strconv.Itoa(root.Val)
	index++
	if root.Left != nil {
		getResult(root.Left, start, d, (en-start)/2+start, index)
	}
	if root.Right != nil {
		getResult(root.Right, (en-start)/2+start+1, d, en, index)
	}
}
