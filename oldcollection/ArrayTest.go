
package main

import "fmt"

func MyArray()  {
	var a [10]int = [10]int{1,2,4,3,5,6,7,8,9}
	fmt.Println(a)

	for i,v:=range a {
		fmt.Println(i, v)
	}

	var b = [...]int{1,2, 3}
	fmt.Println(b)

	var c = new([10]int)
	for i:=1;i<11;i++{
		c[i-1] = i
	}
	fmt.Println(c,*c)



}
/**
生成器的作用:

 */
func main()  {
	MyArray()
}