package main

import (
	"fmt"
	"strings"
)

func main() {
	r := longestValidParentheses(")()())")
	fmt.Println(r)
}

func longestValidParentheses(s string) int {
	if len(s) == 1 || len(s) == 0 {
		return 0
	}

	stack := make([]string, 0)
	ss := strings.Split(s, "")
	for _, v := range ss {
		if len(stack) == 0 {
			if v == ")" {
				continue
			}
			stack = append(stack, v)
			continue
		}

		if stack[len(stack)-1] == "(" && v == ")" {
			stack = append(stack, v)
			continue
		}

		if stack[len(stack)-1] == v {
			stack = stack[:len(stack)]

			continue
		}
		stack = append(stack, v)
	}

	return len(stack)
}

func longestValidParentheses1(s string) int {
	if len(s) == 1 || len(s) == 0 {
		return 0
	}
	if !strings.Contains(s, "()") {
		return 0
	}
	r := findParentheses("()", s)
	return len(r)
}

func findParentheses(find string, src string) string {
	if strings.Contains(src, find) && !strings.Contains(src, find+"()") {
		return find
	}
	if strings.Contains(src, find) {
		find = find + find
		return findParentheses(find, src)
	}

	find = find + find[:len(find)/2]
	return findParentheses(find, src)
}
