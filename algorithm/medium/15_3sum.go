package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3, 0, 0, 0, -1, -1, -1}
	r := threeSum(a)
	fmt.Println(r)
	b := [][]int{0, 0, 0}
}

func threeSum(nums []int) [][]int {
	size := len(nums)
	if size < 3 {
		return nil
	}

	result := [][]int{}

	sort.Ints(nums)

	record := make(map[int]int)
	for _, v := range nums {
		record[v]++
	}

	i := 0
	for ; i < size-2; i++ {
		j := i + 1

		for ; j < size-1; j++ {
			if exist(record, nums[i], nums[j], 0-nums[i]-nums[j]) {
				if replicate(result, []int{nums[i], nums[j], 0 - nums[i] - nums[j]}) {
					continue
				}
				result = append(result, []int{nums[i], nums[j], 0 - nums[i] - nums[j]})
			}
		}
	}

	return result
}

func exist(m map[int]int, a, b, n int) bool {
	for k, _ := range m {
		v := m[k]
		if k == n && v > 0 {
			if k == a {
				v--
			}
			if k == b {
				v--
			}
			if v > 0 {
				v--
				return true
			}
		}
	}

	return false
}

func replicate(s [][]int, f []int) bool {
	for _, v := range s {
		if compare(v, f) {
			return true
		}
	}

	return false
}

func compare(a []int, b []int) bool {
	sort.Ints(a)
	sort.Ints(b)

	for ia, va := range a {
		for ib, vb := range b {
			if ia == ib && va != vb {
				return false
			}
		}
	}

	return true
}
