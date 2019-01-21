package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func main() {
	node := new(Node)
	createTreeManual(node)
	printTree(node)
}

func createTree(node *Node, n int) {
	if n == 0 {
		return
	}

	node.Val = rand.Intn(20)
	n--
	node.Left = new(Node)
	node.Right = new(Node)
	createTree(node.Left, n)
	createTree(node.Right, n)
}

func createTreeManual(node *Node) {
	node.Val = rand.Intn(10)
	node.Left = new(Node)
	node.Right = new(Node)

	node.Left.Val = rand.Intn(90)
	node.Right.Val = rand.Intn(2)
	node.Left.Left = new(Node)
	node.Left.Right = new(Node)
	node.Right.Left = new(Node)

	node.Left.Left.Val = rand.Intn(59)
	node.Left.Right.Val = rand.Intn(100)
	node.Right.Left.Val = rand.Intn(3)

}

func findNodeValue(a, b int, node *Node) int {

}

var r int

func searchNode(a, b int, node *Node) int {

	for node != nil {

	}

	if a == node.Left.Val {
		r = node.Val
		searchNode(a, b, node.Left)
	}

}

func printTree(node *Node) {
	fmt.Println(node.Val)
	if node.Left != nil {
		printTree(node.Left)
	}
	if node.Right != nil {
		printTree(node.Right)
	}
	return
}

func unrecusiveprint(node *Node) {
	for node != nil {
		fmt.Println(node.Val)
	}
}
