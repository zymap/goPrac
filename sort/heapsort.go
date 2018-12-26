package main

import "fmt"

func main() {
	a := []int{2, 5, 8, 4, 6, 2}
	fmt.Println(a)
	heapSortT(a)
	fmt.Println(a)
}

func heapfy(s []int) {
	i := len(s)/2 - 1
	for ; i >= 0; i-- {
		compare(s, i)
	}

}

func compare(s []int, end int) {
	index := end/2 - 1
	for index <= end {
		left := index*2 + 1
		right := index*2 + 2

		if left >= end {
			left = end - 1
		}

		if right >= end {
			right = end - 1
		}

		if s[right] > s[left] && s[right] > s[index] {
			tmp := s[right]
			s[right] = s[index]
			s[index] = tmp
		} else if s[right] < s[left] && s[left] > s[index] {
			tmp := s[left]
			s[left] = s[index]
			s[index] = tmp
		}

		index = index*2 + 1

	}
}

func compareArray(s []int) {

}

func heapSortT(s []int) {
	for i := len(s)/2 - 1; i >= 0; i-- {
		compareIndex(s, len(s), i)
	}

	for i := 0; i < len(s); i++ {
		tmp := s[len(s)-1-i]
		s[len(s)-1-i] = s[i]
		s[i] = tmp

		for j := (len(s)-1-i)/2 - 1; j >= 0; j-- {
			compareIndex(s, len(s)-i-1, j)
		}

	}
}

func compareIndex(s []int, length int, index int) {
	for index <= length/2-1 {
		left := index*2 + 1
		right := index*2 + 2

		if left >= length {
			left = length - 1
		}

		if right >= length {
			right = length - 1
		}

		if s[right] > s[left] && s[right] > s[index] {
			tmp := s[right]
			s[right] = s[index]
			s[index] = tmp
		} else if s[right] < s[left] && s[left] > s[index] {
			tmp := s[left]
			s[left] = s[index]
			s[index] = tmp
		}

		index = index*2 + 1
	}
}

func swap(a, b int) (int, int) {
	return b, a
}
