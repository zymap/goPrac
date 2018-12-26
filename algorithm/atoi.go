package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	i := myAtoi("                     -99999                ")
	fmt.Println(i)

	s := "hello"
	for _, v := range s {
		fmt.Println(string(v))
	}
}

func myAtoi(str string) int {
	news := strings.TrimSpace(str)

	tmp := ""
	flag := false

	for _, c := range news {
		n, err := strconv.Atoi(string(c))
		if err != nil {
			if (string(c) == "-" || string(c) == "+") && !flag {
				tmp += string(c)
				flag = true
				continue
			}
			break
		}

		tmp += strconv.Itoa(n)
		flag = true
	}

	i, err := strconv.Atoi(tmp)
	if err != nil {
		fmt.Println(err.Error())
	}

	if i > math.MaxInt32 {
		return math.MaxInt32
	}

	if i < math.MinInt32 {
		return math.MinInt32
	}

	return i
}
