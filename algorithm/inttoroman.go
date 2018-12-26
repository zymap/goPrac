package main

import "fmt"

func main() {
	r := intToRoman(58)
	fmt.Println(r)

}

var romanMap = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

func intToRoman(num int) string {
	var res string

	var thousands int
	if num >= 1000 {
		thousands = num / 1000
		for i := thousands; i > 0; i-- {
			res += "M"
		}
	}

	var hundrund int
	if num >= 100 {
		num = (num - thousands*1000)
		hundrund = num / 100
		for i := hundrund; i > 0; i-- {
			if i == 9 {
				res += "CM"
				break
			}

			if i == 5 {
				res += "D"
				break
			}

			if i > 5 {
				res += "D"
				i -= 4
				continue
			}

			if i == 4 {
				res += "CD"
				break
			}

			res += "C"
		}
	}

	var ten int
	if num >= 10 {
		num = num - hundrund*100
		ten = num / 10
		for i := ten; i > 0; i-- {
			if i == 9 {
				res += "XC"
				break
			}

			if i == 5 {
				res += "L"
				break
			}

			if i > 5 {
				res += "L"
				i -= 4
				continue
			}

			if i == 4 {
				res += "XL"
				break
			}

			res += "X"
		}
	}
	n := num - ten*10
	for ; n > 0; n-- {
		if n == 9 {
			res += "IX"
			break
		}

		if n == 5 {
			res += "V"
			break
		}

		if n > 5 {
			res += "V"
			n -= 4
			continue
		}

		if n == 4 {
			res += "IV"
			break
		}

		res += "I"
	}

	return res
}
