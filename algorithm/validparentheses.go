package main

import (
	"fmt"
	"strings"
)

func main() {
	r := isVlid("{[]}")
	fmt.Println(r)
}

func isVlid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	if string(s[0]) == "}" || string(s[0]) == ")" || string(s[0]) == "]" {
		return false
	}

	mysatck := New()
	brackets := strings.Split(s, "")
	for i, v := range brackets {
		if !mysatck.null() && isRight(v) {
			if check(brackets[i-1], v) {
				mysatck.pop()
				continue
			}
		}
		mysatck.push(v)
	}

	return mysatck.null()
}

func isRight(s string) bool {
	if s == ")" || s == "]" || s == "}" {
		return true
	}

	return false
}

type stack struct {
	Val   []string
	Index int
}

func New() *stack {
	value := []string{}
	return &stack{Val: value, Index: 0}
}

func (s *stack) push(braket string) *stack {
	s.Val = append(s.Val, braket)
	s.Index++
	return s
}

func (s *stack) pop() string {
	r := s.Val[s.Index-1]
	s.Index--
	return r
}

func (s *stack) null() bool {
	return s.Index == 0
}

func check(left, right string) bool {
	if left == "[" && right == "]" {
		return true
	}

	if left == "{" && right == "}" {
		return true
	}

	if left == "(" && right == ")" {
		return true
	}

	return false
}
