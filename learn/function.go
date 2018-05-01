package learn

import "fmt"

func FunctionX(a *int)  {
	b := a
	fmt.Println("function X :")
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(*b)
}

func FunctionY(a int) {
	b := &a
	fmt.Println("function Y :")
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(*b)
}