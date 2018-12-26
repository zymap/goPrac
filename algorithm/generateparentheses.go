package main

func main() {

}

func generateParenthesis(n int) []string {
	r := []string{}
	if n == 0 {
		return r
	}

	leftBrackets := "("
	rightBrackets := ")"

	for i := 0; i < n; i++ {
		r = append(r, leftBrackets)

	}

}
