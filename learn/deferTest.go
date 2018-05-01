package learn

import (
	"fmt"
)

func MyDefer()  {
	fmt.Println("this is begining...")
	for i := 0; i < 10; i++ {
		defer test(i)
	}

	fmt.Println("this is ending...")
}

func test(i int)  {
	fmt.Println("hello i m defer...",i)
}