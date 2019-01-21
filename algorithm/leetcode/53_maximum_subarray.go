package main

func main() {

}

func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	max := dp[0]

	maxFunc := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < n; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}

		max = maxFunc(max, dp[i])
	}

	return max
}
