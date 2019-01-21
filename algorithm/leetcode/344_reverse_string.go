package main

import "fmt"

func main() {
	str := "hello"
	r := reverseString(str)
	fmt.Print(r)

}

func reverseString(s string) string {
	ss := []rune(s)
	for i := 0; i < len(ss)/2; i++ {
		temp := ss[len(ss)-i-1]
		ss[len(ss)-i-1] = ss[i]
		ss[i] = temp
	}

	return string(ss)
}
