package main

import "fmt"

func main() {
	fmt.Println(1 << 31)
	s := []int{4, 3, 2, 6}
	r := plan(s)
	fmt.Println(r)
}

func maxRotateFunction(A []int) int {
	if len(A) == 0 {
		return 0
	}

	max := -1 << 31

	for i := 0; i < len(A); i++ {
		var s []int

		s = append(s, A[len(A)-i:]...)
		s = append(s, A[:len(A)-i]...)
		r := compute(s)
		if r > max {
			max = r
		} else {
			max = max
		}
	}

	return max
}

func compute(s []int) int {
	result := 0

	for i, v := range s {
		if v == -1<<31 {
			result = v
			continue
		}
		result += i * v
	}

	return result
}

func plan(A []int) int {
	sum := 0
	val := 0

	for i, v := range A {
		sum += v
		val += i * v
	}

	res := val
	for i := len(A) - 1; i >= 1; i-- {
		val = val + sum - A[i]*len(A)
		if val > res {
			res = val
		}
	}

	return res
}
