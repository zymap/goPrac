package main

import "fmt"

type B struct {
	Name string
}

func (b *B) Func1(name string) {
	b.Name = name
}

func (b B) Func2(name string) {
	b.Name = name
}

func main() {
	b := new(B)
	b.Func1("aaa")
	fmt.Println(b)

	b.Func2("bbb")
	fmt.Println(b)
}
