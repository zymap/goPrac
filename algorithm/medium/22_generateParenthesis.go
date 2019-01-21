package main

import "fmt"

func main() {
	r := generateParenthesis(3)
	fmt.Println(r)
}

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	combine(&result, "", 0, 0, n)
	return result
}

func combine(result *[]string, cur string, open, close, max int) {
	if len(cur) == 2*max {
		*result = append(*result, cur)
		return
	}

	if open < max {
		combine(result, cur+"(", open+1, close, max)
	}
	if close < open {
		combine(result, cur+")", open, close+1, max)
	}
}
