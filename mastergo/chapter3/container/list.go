package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func printList(l *list.List) {
	for t := l.Back(); t != nil; t = t.Prev() {
		fmt.Print(t.Value, "\t")
	}
	fmt.Println()
	for t := l.Front(); t != nil; t = t.Next() {
		fmt.Print(t.Value, "\t")
	}
	fmt.Println()
}

func main() {
	values := list.New()
	e1 := values.PushBack("1111")
	e2 := values.PushBack("2222")

	values.PushFront("3333")

	values.InsertBefore("4444", e1)
	values.InsertAfter("5555", e2)
	values.Remove(e2)
	values.Remove(e2)

	values.InsertAfter("55555555", e2)
	values.PushBackList(values)

	printList(values)

	values.Init()
	fmt.Printf("after init(): %v\n", values)

	for i := 0; i < 20; i++ {
		values.PushFront(strconv.Itoa(i))
	}
	printList(values)
}
