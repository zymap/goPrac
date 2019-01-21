package main

import "fmt"

/*
   Dynamic programming.
    Best Time to Buy and Sell Stock.
*/

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	min := prices[0]
	max := 0
	for _, v := range prices {
		if v > min {
			if max < v-min {
				max = v - min
			}
		} else {
			min = v
		}
	}

	return max
}

func findPeaksAndValleys() {

}

func main() {
	stock := []int{7, 1, 5, 3, 6, 4}
	r := maxProfit(stock)
	fmt.Print(r)
}
