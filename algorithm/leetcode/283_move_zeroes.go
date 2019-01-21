package main

import "sort"

func main() {

}

func moveZeroes(nums []int) {
	if len(nums) == 1 {
		if nums[0] == 0 {
			return
		}
	}

	j := len(nums) - 1
	for i := 0; i < j && j >= 0; i++ {
		if nums[i] == 0 {
			nums[i] = nums[j]
			nums[j] = 0
			j--
		}
	}

	sort.Ints(nums[:j])
}
