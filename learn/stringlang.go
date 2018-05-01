package learn

import "fmt"

func StringUtil()  {
	str := "this is my test string."
	str = "abcdefg"
	char := []rune(str)
	pos := 0
	for pos < len(char) {
		fmt.Println(char[pos])
		pos++
	}

	for i,ch := range char{
		fmt.Println(i,ch)
	}

	for i,ch := range str{
		fmt.Println(i,ch)
	}
}