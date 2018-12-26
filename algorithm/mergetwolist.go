package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//type ListNode struct {
//    Val int
//    Next *ListNode
//}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	l := new(ListNode)
	r := l

	for l1 != nil || l2 != nil {
		if l1 == nil {
			*l = *l2
			break
		}

		if l2 == nil {
			*l = *l1
			break
		}

		if l1.Val <= l2.Val {
			l.Val = l1.Val
			l1 = l1.Next
		} else {
			l.Val = l2.Val
			l2 = l2.Next
		}

		l.Next = new(ListNode)
		l = l.Next
	}

	return r
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l := createList(2)
	printList(l)
	l1 := createList(3)
	printList(l1)

	fmt.Println("========================")
	r := mergeTwoLists(l, l1)
	printList(r)

	//r := mergeTwoLists(l, l1)
	//printList(r)
}

func createList(n int) *ListNode {
	l := new(ListNode)
	r := l
	for i := 1; i <= n-1; i++ {
		l.Val = i
		l.Next = new(ListNode)
		l = l.Next
	}

	l.Val = n

	return r
}

func printList(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}
