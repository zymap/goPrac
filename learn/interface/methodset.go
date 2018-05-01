package _interface

import "fmt"

type List []int

func (this *List) Len() int {
	return len(*this)
}

func (this *List) Append(value int)  {
	*this = append(*this,value)
}

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int)  {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len() * 10 > 42
}

func MethodSet()  {
	//list := List{1, 2}
	//CountInto(list, 1, 6)
	//if LongEnough(list) {
	//	fmt.Println("list is long enougn")
	//}

	ptr_list := new(List)
	CountInto(ptr_list, 1, 6)
	if LongEnough(ptr_list) {
		fmt.Println("ptr_list is long enough")
	}
}