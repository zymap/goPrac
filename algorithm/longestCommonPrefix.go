package main

import "fmt"

func main() {

	inputs := []string{
		"c",
		"acc",
		"ccc",
	}

	r := longestCommonPre(inputs)
	fmt.Println(r)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	var table string
	for _, v := range strs[1:] {

	}
}

func longestCommonPre(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}
	var table string
	flag := false
	for _, v := range strs[1:] {
		result := find(strs[0], v)

		if len(result) < len(table) || (table == "" && !flag) {
			table = result
			flag = true
		}
	}
	return table
}

func find(str1, str2 string) string {
	var res string

	l := len(str1)
	if len(str1) > len(str2) {
		l = len(str2)
	}

	for i := 0; i < l; i++ {
		if str1[i] != str2[i] {
			break
		}
		res += string(str1[i])
	}

	return res
}
