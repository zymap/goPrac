package main

import (
	"container/heap"
	"fmt"
)

type heapInt []int

func (n *heapInt) Pop() interface{} {
	old := *n
	x := old[len(old)-1]
	new := old[0 : len(old)-1]
	*n = new
	return x
}

func (n *heapInt) Push(x interface{}) {
	*n = append(*n, x.(int))
}

func (n heapInt) Len() int {
	return len(n)
}

func (n heapInt) Less(a, b int) bool {
	return n[a] < n[b]
}

func (n heapInt) Swap(a, b int) {
	n[a], n[b] = n[b], n[a]
}

func main() {
	myHeap := &heapInt{1, 2, 3, 4}
	heap.Init(myHeap)
	size := len(*myHeap)
	fmt.Println("Heap Size : ", size)
	fmt.Println(myHeap)

	for i := 0; i < 4; {
		myHeap.Push(i)
		i++
	}
	fmt.Println("Heap Size : ", len(*myHeap))
	fmt.Println(myHeap)

	heap.Init(myHeap)
	fmt.Println(myHeap)
}
