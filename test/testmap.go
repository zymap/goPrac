package main

import (
	"fmt"
	"strconv"
)

type Tmp struct {
	Id   int
	name string
}

/*
*	This case used to verify the map. When use the map, copy it, than use the copier of the map values.
*	Otherwise, the map will be changed.
 */

func main() {
	m := make(map[string]*Tmp)

	for i := 0; i < 10; i++ {
		id := i
		name := "a" + strconv.Itoa(i)
		tmp := Tmp{
			id,
			name,
		}
		m[strconv.Itoa(i)] = &tmp
	}
	fmt.Println("create map : ")
	fmt.Println(m)
	printMap(m)

	fmt.Println("use map: ")
	for i, v := range m {
		if i == "3" {
			v.name = "hello"
		}
	}

	fmt.Println("after use map:")
	fmt.Println(m)
	printMap(m)

	fmt.Println("use again: ")
	for i, v := range m {
		value := v
		if i == "3" {
			value.name = "ha"
		}
	}
	fmt.Println("after use again: ")
	fmt.Println(m)
	printMap(m)

}

func printMap(m map[string]*Tmp) {
	for i, v := range m {
		fmt.Printf("key %s ,\t value %v \n", i, *v)
	}
	fmt.Println("========================")
}
