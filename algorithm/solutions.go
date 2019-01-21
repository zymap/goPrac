package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	s := "1[a2[bc]2[bc]b]"
	r := solutions(s)
	fmt.Println(r)
}

func solutions(s string) string {
	result := ""

	left := []int{}
	right := []int{}

	for i, v := range strings.Split(s, "") {
		if v == "[" {
			left = append(left, i)
		}
		if v == "]" {
			right = append(right, i)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(right)))

	for i, v := range left {
		num, _ := strconv.Atoi(string(s[v-1]))
		str := s[v+1 : right[i]]
		result = strings.Replace(s, str, multi(num, str), 1)
	}

	result = strings.Replace(result, "[", "", -1)
	result = strings.Replace(result, "]", "", -1)

	return result
}

func recursiveMulti(s string) string {
	f := 0
	record := []int{}
	for i, v := range s {
		if string(v) == "[" {
			record = append(record, i)
			f++
		}

		if string(v) == "]" {
			f--
			if f == 0 {
				r := multi(s[i])
			}
		}
	}
}

func multi(n int, s string) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}
