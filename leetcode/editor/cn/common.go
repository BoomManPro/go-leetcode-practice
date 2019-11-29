package leetcode_cn

import (
	"bytes"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (curr *ListNode) String() string {
	var buf bytes.Buffer
	var temp *ListNode = curr
	for temp != nil {
		buf.WriteString(strconv.Itoa(temp.Val))
		if temp.Next != nil {
			buf.WriteString("->")
		}
		temp = temp.Next
	}
	return buf.String()
}

func (curr *ListNode) addNode(node *ListNode) *ListNode {
	curr.Next = node
	return node
}

func (curr *ListNode) addNodeForValue(value int) *ListNode {

	return curr.addNode(NewListNode(value))
}

func NewListNode(value int) *ListNode {
	return &ListNode{
		Val:  value,
		Next: nil,
	}
}
