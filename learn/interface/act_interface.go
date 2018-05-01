package _interface

import "fmt"

type IDuck interface {
	Quack()
	Walk()
}

func DuckDance(duck IDuck)  {
	for i := 1; i <= 1; i++ {
		duck.Quack()
		duck.Walk()
	}
}

type Bird struct {

}

func (this *Bird) Quack()  {
	fmt.Println("i am quacking bird")
}

func (this *Bird) Walk()  {
	fmt.Println("i am walking bird")
}

func Act()  {
	b := new(Bird)
	//fmt.Println(b)
	DuckDance(b)
}