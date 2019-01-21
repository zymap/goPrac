package main

import "fmt"

func main() {
	str := "a"
	i := 1
	r := longestSubstring(str, i)
	fmt.Println(r)
}

func longestSubstring(s string, k int) int {

	res, i, n := 0, 0, len(s)

	for i+k <= n {
		m := make([]int, 26)
		mask := 0
		max_idx := i

		for j := i; j < n; j++ {
			t := s[j] - 'a'
			m[t]++

			if m[t] < k {
				mask |= 1 << t
			} else {
				mask &= (^(1 << t))
			}

			if mask == 0 {
				if res >= j-i+1 {
					res = res
				} else {
					res = j - i + 1
				}
				max_idx = j
			}
		}
		i = max_idx + 1
	}
	return res
}

func longestSubString1(s string, k int) int {
	return longestSub(s, k, 0, len(s)-1)
}

func longestSub(s string, k, start, end int) int {
	if start > end {
		return 0
	}

	count := make([]int, 26)

	for _, v := range s {
		count[v-'a']++
	}

	for i := 0; i < 26; i++ {
		pos := strings.index
	}
}
