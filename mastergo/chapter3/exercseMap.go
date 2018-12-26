package main

import "fmt"

func main() {
	iMap := make(map[string]int)
	fmt.Println(iMap)
	anotherMap := map[string]int{
		"k1": 12,
		"k2": 13,
	}

	fmt.Println("access")
	fmt.Println(anotherMap["k1"])

	fmt.Println("delete")
	delete(anotherMap, "k1")

	for key, value := range anotherMap {
		fmt.Println(key, value)
	}
}
