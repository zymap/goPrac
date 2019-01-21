package main

import "fmt"

func main() {
	nums := []int{1}
	target := 0
	r := searchRange(nums, target)
	fmt.Println(r)
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	start := -1
	tail := -1

	begin := 0
	end := len(nums) - 1

	for begin <= end {
		median := (begin + end) / 2
		if nums[median] > target {
			end--
		} else if nums[median] < target {
			begin++
		} else {
			if nums[begin] < target {
				begin++
			}
			if nums[end] > target {
				end--
			}
			if nums[begin] == nums[end] {
				return []int{begin, end}
			}
		}
	}

	return []int{start, tail}
}
