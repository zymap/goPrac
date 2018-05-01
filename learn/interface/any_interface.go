package _interface

import "fmt"

type Element interface {}

type Vector struct {
	value [10]Element
}

func (this *Vector) Set(index int,value Element) {
	this.value[index] = value
}

func (this *Vector) Get(index int) Element {
	return this.value[index]
}

func MyVector() {
	v := Vector{}

	v.Set(1, 0)
	v.Set(2, "aaa")

	fmt.Println("the v[1] is : ", v.Get(1))
	fmt.Println("the v[2] is : ", v.Get(2))
}

func CopyToNilInterface()  {
	dataslice := []int{1, 2, 3, 4, 5}
	interfaceslice := make([]interface{}, 5)
	for i, v := range dataslice {
		interfaceslice[i] = v
	}

	fmt.Println("interface slice is : ")
	for _, v := range interfaceslice {
		fmt.Print(v, "\t")
	}
	fmt.Println()
}