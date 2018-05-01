package datastructure

import (
	"fmt"
	//"unicode"
)


var first *Linkedlist
var end *Linkedlist

type Linkedlist struct {
	value interface{}
	pre *Linkedlist
	next *Linkedlist
}

func (list *Linkedlist)add(a interface{}) (ok bool)  {
	list.linklast(a)
	return true
}

func (list *Linkedlist) linklast(a interface{}){
	var node Linkedlist
	node.value = a
	if first == nil {
		first = &node
		end = &node
		return
	}
	node.pre = list
	list.next = &node
	list = &node
}

func (list *Linkedlist) getLast() (result interface{})  {
	result = end.value
	return result
}

func (list *Linkedlist) print() {
	fmt.Println("hello")
	fmt.Println(first.value)
	var node = first
	fmt.Println(node.value)
	for node.next != nil {
		fmt.Println("1")
		fmt.Println(node.value)
		node = node.next
	}
}

func Test() {
	var s = struct {
		a int
	}{1}
	var list= Linkedlist{}
	list.add(s)
	list.add("aaaa")
	list.print()
	//fmt.Println(list.getLast())
	//fmt.Println(list.getLast())
}