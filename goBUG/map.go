package main

import "fmt"

func main() {
	//m := make(map[string]string)

	//tests := []struct {
	//	a int
	//	b string
	//}{
	//	{1, "aaa"},
	//	{2, "bbb"},
	//}
	//
	//var t []*struct {
	//	a int
	//	b string
	//}
	//
	//for _, v := range tests {
	//	println(v.a, "\t", v.b)
	//}
	//
	//for _, a := range tests {
	//	append(t, &a)
	//}
	//
	//for _, v := range tests {
	//	println(v.a, "\t", v.b)
	//}

	bug()
}

func bug() {

	a := 1
	b := 2

	tests := []struct {
		i *int
	}{
		{&a},
		{&b},
	}

	for _, v := range tests {
		fmt.Println(*v.i)
		*v.i = 10
	}

	for _, v := range tests {
		fmt.Println(*v.i)
	}

	for i, v := range tests {
		v = *v
		v.i = i
	}

	for _, v := range tests {
		fmt.Println(*v.i)
	}
}
