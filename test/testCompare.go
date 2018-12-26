package main

import (
	"fmt"
	"reflect"
)

type store struct {
	key   int
	value string
}

type Test struct {
	Content []*store
}

type Test1 struct {
	TT struct {
		Content []*store
	}
}

func main() {
	store1 := store{
		1,
		"aaa",
	}

	store2 := store{
		1,
		"aaa",
	}

	if !reflect.DeepEqual(store1, store2) {
		fmt.Println("store1 and store2 not equal.")
	}

	t1 := Test{
		[]*store{&store1, &store2},
	}

	t2 := Test{
		[]*store{&store1, &store2},
	}

	if !reflect.DeepEqual(t1, t2) {
		fmt.Println("t1 and t2 not equal.")
	}

	t3 := Test1{
		TT: struct{ Content []*store }{Content: []*store{&store1}},
	}

	t4 := Test1{
		TT: struct{ Content []*store }{Content: []*store{&store2}},
	}

	if !reflect.DeepEqual(t3, t4) {
		fmt.Println("t3 and t4 not equal.")
	}

	fmt.Println("over.")
}
