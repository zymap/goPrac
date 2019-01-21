package main

import (
	"fmt"
	"strconv"
)

func main() {
	r := fizzBuzz(10)
	fmt.Print(r)
}

func fizzBuzz(n int) []string {
	result := make([]string, 0)

	if n == 0 {
		return result
	}

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "FizzBuzz")
			continue
		}

		if i%3 == 0 {
			result = append(result, "Fizz")
			continue
		}

		if i%5 == 0 {
			result = append(result, "Buzz")
			continue
		}

		result = append(result, strconv.Itoa(i))
	}

	return result
}

func fizzBuzz1(n int) []string {
	result := make([]string, n)
	for i := 1; i <= n; i++ {
		result[i-1] = strconv.Itoa(i)
	}

	for i := 3; i <= n; i += 3 {
		result[i-1] = "Fizz"
	}

	for i := 5; i <= n; i += 5 {
		if result[i-1] == "Fizz" {
			result[i-1] = "FizzBuzz"
			continue
		}
		result[i-1] = "Buzz"
	}

	return result
}
