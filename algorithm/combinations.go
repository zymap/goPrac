package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := letterCombinations("234")
	fmt.Println(a)
}

var table = map[int][]string{

	2: []string{"a", "b", "c"},
	3: []string{"d", "e", "f"},
	4: []string{"g", "h", "i"},
	5: []string{"j", "k", "l"},
	6: []string{"m", "n", "o"},
	7: []string{"p", "q", "r", "s"},
	8: []string{"t", "u", "v"},
	9: []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	var a []string

	r := strings.Split(digits, "")
	for _, v := range r {
		i, _ := strconv.Atoi(v)
		a = getCombinations(a, table[i])
	}

	return a
}

func getCombinations(strs1, strs2 []string) []string {
	if len(strs1) == 0 && len(strs2) != 0 {
		return strs2
	}
	var res []string
	for _, v := range strs1 {
		for _, s := range strs2 {
			res = append(res, v+s)
		}
	}
	return res
}
