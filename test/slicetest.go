package main

import "fmt"

// slice is a 前闭后开区间
func test() {
	s := make([]int, 10)

	// init
	i := 0
	for i < 10 {
		s[i] = i
		i++
	}

	// 1 -- 5
	fmt.Println(s[:5])

	fmt.Println(s[5])

	// 5 -- 10
	fmt.Println(s[5:6])
}

func main() {
	createAndShow()
}

// 对于slice
func createAndShow() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	modify(a)
	fmt.Println(a)
	modify2(a)
	fmt.Println(a)
	modify3(a)
	fmt.Println(a)
	modify4(a)
	fmt.Println(a)
	modify5(a)
	fmt.Println(a)
}

// 直接操作slice
func modify(a []int) {
	for i, v := range a {
		a[i] = v + 1
	}
}

// 将slice赋值到新的变量，但是操作的还是同一个slice， 操作同一个地址
func modify2(a []int) {
	b := a
	for i, v := range b {
		b[i] = v + 1
	}
}

// 在新的域内新定义slice的指向
func modify3(a []int) {
	b := []int{2, 3, 4, 5, 6, 6}
	a = b
}

func modify4(a []int) {
	copy(a, a[:len(a)/2])
}

func modify5(a []int) {
	b := append(a, 100)
	b[1] = 10
}
