package main

import (
	"fmt"
)

func main() {
	l := new(ListNode)
	lHead := l
	l.Val = 1
	l.Next = new(ListNode)
	l = l.Next
	l.Val = 2
	l.Next = new(ListNode)
	l = l.Next
	l.Val = 3
	l.Next = new(ListNode)
	l = l.Next
	l.Val = 4

	r := removeNthFromEnd(lHead, 2)
	fmt.Println(r.Val)
	printLists(r)
}

func printLists(h *ListNode) {
	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	index := 0
	nindex := n
	result := new(ListNode)
	resultHead := result

	for head != nil {

		if index-nindex == n && head.Next == nil {
			head = head.Next
			index++
			continue
		}

	}

	for head != nil {
		if index == n {
			head = head.Next
			index++
			continue
		}
		result.Val = head.Val
		if head.Next != nil {
			result.Next = new(ListNode)
			result = result.Next
		}

		head = head.Next
		index++
	}

	return resultHead
}
