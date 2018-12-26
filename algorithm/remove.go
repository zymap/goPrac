package main

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	a := head
	b := head

	i := 0
	for i < n {
		a = a.Next
	}

	for a != nil {
		a = a.Next
		b = b.Next
	}

	b.Next = b.Next.Next

	return b
}
