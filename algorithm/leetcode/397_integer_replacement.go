package main

import "fmt"

func main() {
	r := compare(65535)
	fmt.Println(r)
}

func integerReplacement(n int) int {
	step := 0

	if n == 1 {
		return step
	}

	num := n
	if num%2 == 1 {
		num = num - 1
		step++
	}
	for {
		if num == 1 {
			break
		}
		num = num / 2
		step++
		if num == 1 {
			break
		}
		if num%2 == 1 {
			num = num - 1
			step++
		}
	}

	return step
}

func compare(n int) int {
	step := 0
	if n == 1 {
		return step
	}

	num := n
	for num%2 == 0 && num != 1 {
		step++
		num = num / 2
	}

	if num%2 == 1 && num != 1 {
		s1 := compare(num + 1)
		s2 := compare(num - 1)
		if s1 > s2 {
			step = step + s2 + 1
		} else {
			step = step + s1 + 1
		}
	}

	return step
}
