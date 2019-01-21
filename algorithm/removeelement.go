package main

import "fmt"

func main() {
	l := []int{1, 2, 3, 1, 1, 1}
	r := removeElement(l, 1)
	fmt.Println(r)
	fmt.Println(l)
}

func removeElement(nums []int, val int) int {
	j := 0
	c := 0
	for _, v := range nums {
		if v == val {
			continue
		}

		if c == 0 {
			nums[j] = v
			j++
			c++
			continue
		}

		nums[j] = v
		j++
		c++
	}

	return c
}
