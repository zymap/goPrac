package main

import (
	"fmt"
	"strings"
)

func main() {
	r := strStr("mississippi", "issip")
	fmt.Println(r)
}

func strStr(haystack string, needle string) int {
	if !strings.Contains(haystack, needle) {
		return -1
	}

	if needle == "" {
		return 0
	}

out:
	for i, v := range haystack {
		if string(v) == string(needle[0]) {
			for ia, iv := range needle {
				if string(iv) != string(haystack[i+ia]) {
					continue out
				}
			}
			return i
		}
	}

	return -1
}
