package main

import (
	"fmt"
	"reflect"
	"sync"
)

func main() {
	var smap sync.Map

	smap.Store("1", 1)

	smap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		fmt.Println(reflect.TypeOf(key))
		fmt.Println(key.(string))
		var a string
		a = key.(string)
		fmt.Println(a)
		fmt.Println(reflect.TypeOf(value))
		return true
	})
}
