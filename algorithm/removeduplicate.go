package main

import "fmt"

func main() {
	l := []int{1, 1, 2}
	r := removeDuplicates1(l)
	fmt.Println(l)
	fmt.Println(r)
}

func removeDuplicates(nums []int) int {
	store := make(map[int]bool)

	c := 0
	j := 0
	for i, v := range nums {

		if !store[v] {
			tmp := nums[j]
			nums[j] = nums[i]
			nums[i] = tmp
			j++
			store[v] = true
			c++
			continue
		}

	}
	return c
}

func removeDuplicates1(nums []int) int {
	var t int
	j := 0
	for i, v := range nums {
		if t != v {
			t = v
			tmp := nums[j]
			nums[j] = nums[i]
			nums[i] = tmp
			j++
		}
	}
	return j
}
