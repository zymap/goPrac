package main

import "fmt"

type listNode struct {
	Val  int
	Next *listNode
}

func swapPairs(head *listNode) *listNode {
	if head == nil {
		return head
	}

	if head.Next == nil {
		return head
	}

	l := new(listNode)
	r := l
	for head != nil && head.Next != nil {
		if l == nil {
			l = new(listNode)
		}
		l.Next = new(listNode)
		l.Val = head.Next.Val
		l = l.Next
		l.Val = head.Val

		l = l.Next

		head = head.Next.Next
	}

	return r

}

func createTest(n int) *listNode {
	l := new(listNode)
	r := l

	for ; n > 0; n-- {
		l.Val = n
		l.Next = new(listNode)
		l = l.Next
	}

	return r
}

func printList(l *listNode) {
	for l != nil {
		fmt.Print(l.Val, "\t")
		l = l.Next
	}
	fmt.Println()
}

func TestSwapNodesInPairs() {
	l := createTest(4)
	printList(l)
	r := swapPairs(l)
	printList(r)
}
