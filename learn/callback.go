package learn

import "fmt"

func MainFunc()  {
	fmt.Println(compute(1,2,add))
}

func compute(i int, j int, f func(int,int)(int))(result int)  {
	result = f(i,j)
	return
}

func add(i int, j int)(result int)  {
	result = i + j
	return
}