package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	if k < 0 {
		return head
	}

	i := 0

	l := new(ListNode)
	r := l

	record := make([]int, k)

	for head != nil {
		if i == k {
			break
		}
		record = append(record, head.Val)
		head = head.Next
		i++
	}

	for _, v := range record {
		l.Val = v
		l.Next = new(ListNode)
		l = l.Next
	}

	for head != nil {
		l.Val = head.Val
		head = head.Next
		l.Next = new(ListNode)
		l = l.Next
	}

	return r

}
