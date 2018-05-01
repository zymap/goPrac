package collection

import (
	system "fmt"
	"bytes"
	"sort"
)

func Array()  {
	var a [10]int= [10]int{1,22,3,5,6, 8}
	system.Println(a)
	//fmt.Println(a)
	for i,v:=range a {
		system.Println(i,v)
	}
}


/**
new 出来的数组是一个引用类型  而声明的变量只是将a的引用拷贝 并不会拷贝数组中的值，两个修改还是同一个地方
需要将值也拷贝出来
 */
func Array1()  {
	a := new([3]int)
	a[0] = 10
	a[1] = 20
	system.Println(a)

	b := a
	b[1] = 99
	system.Println(a)
	system.Println(b)

	c := *a
	c[0] = 1100
	system.Println(a)
	system.Println(c)
}

func Array11()  {
	a := [6]int{1, 2, 3, 4, 5, 6}
	system.Println("a : ", a)

	b := a
	system.Println("b : ", b)

	b[0] = 100
	system.Println("update b[0] : ",b)
	system.Println("after update a : ", a)

	system.Println("=========split========")

	aa := []int{1, 2, 3, 4, 5, 6}
	system.Println("aa : ", aa)

	bb := aa
	system.Println("bb : ", bb)

	bb[0] = 100
	system.Println("update bb[0] : ",bb)
	system.Println("after update aa : ", aa)
}

func Array2()  {
	a := [...]string{"1","2","3","4","5"}
	for i := range a{
		system.Println(a[i])
	}
}

func Array3()  {
	array := []int{1, 2, 3, 4, 5}
	//var array [5]int= [5]int{1, 2, 3, 4, 5}
	system.Println("before add ", array)
	//addArray(array)
	system.Println("after add ", array)

	brray := array
	brray[0] = 100
	system.Println("brray:", brray)
	system.Println("array:", array)
	system.Println("===============")
	var bb [3]int = [3]int{1, 2, 3}
	system.Println(bb)
	cc := bb
	system.Println(cc)
	cc[0] = 100
	system.Println(bb)
	system.Println(cc)
	system.Println("===============")
	dd := 20
	b := dd
	b = 30
	system.Println(b)
	system.Println(dd)

	var xx = [5]int{0,0,0,0,0}
	a := new([5]int)
	a = &xx
	system.Println(a)
	//*a = array
	system.Println(a)
	aa := a
	system.Println(aa)
	(*aa)[0] = 100
	system.Println(aa)
	system.Println("before add ", a)
	//addArray(a)
	system.Println("after add ", a)
}

func Array4()  {
	array := new([3]int)
	array[0] = 1
	array[1] = 2
	array[2] = 3
	system.Println("array : ",array)
	a := array
	a[0] = 100
	system.Println("a : ",a)
	system.Println("array : ",array)
}

func addArray(array []int)  {
	for i := range array {
		array[i] += 1
	}
}

func Array5()  {
	array := [5]int{0: 1, 1: 2, 4: 10}
	system.Println(array)
}

func Slice1()  {
	var s []int
	slice := []int{}
	system.Println("slice s : " ,s)
	system.Println("slice slice : ", slice)
	system.Println("slice slice cap : ", cap(slice))
	system.Println("slic slice len ： ", len(slice))
}

func Slice2()  {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[2:]
	system.Println("slice : ", slice)
	system.Println("slice len : ", len(slice))
	system.Println("slice cap : ", cap(slice))
}

func mes(mes string, slice []int)  {
	system.Println(mes, " : ", slice)
	system.Println(mes, " len : ", len(slice))
	system.Println(mes, " cap : ", cap(slice))
}

func Slice3()  {
	slice := make([]int, 3)
	mes("slice", slice)

	slice1 := make([]int, 3, 5)
	mes("slice1", slice1)

	slice3 := new([20]int)[2:13]
	mes("slice2", slice3)

	slice4 := new([20]int)[3:15]
	mes("slice4", slice4)
}

func Slice4()  {
	slice := make([]int, 3, 5)
	slice[2] = 10
	mes("slice", slice)
}

func Slice5()  {
	slice := []int{1, 2, 3, 4, 5}
	mes("slice", slice)

	for i,v := range slice {
		system.Println("slice i : ", i, " value : ", v)
	}

	//update1
	for _,v := range slice {
		v += 2
	}

	for i, v := range slice {
		system.Println("update1 slice i : ", i, " value : ", v)
	}

	//update2
	for i, _ := range slice {
		slice[i] +=2
	}

	for i, v := range slice {
		system.Println("update2 slice i : ", i, " value : ", v)
	}
}

func Reslice1()  {
	//make slice len + 1
	slice1 := make([]int, 2, 10)
	mes("slice1", slice1)

	slice1 = slice1[:len(slice1)+1]
	mes("resize slice1", slice1)
}

func Reslice2()  {
	slice := new([10]int)[ 1 : 1 ]
	mes("slice", slice)
	slice = slice[ 1 : 2 ]
	mes("slice", slice)
}

func CopySlice()  {
	slice_from := []int{1, 2, 3}
	slice_to := make([]int,5, 5)
	mes("slice_from", slice_from)
	mes("slice_to", slice_to)
	n := copy(slice_to, slice_from)
	system.Println("copy numbers : ", n)
	mes("copied slice_from", slice_from)
	mes("copied slice_to", slice_to)
}

func SortSlice() {
	slice := []int{2, 6, 1, 8, 4, 3, 90}
	mes("before sorted : ", slice)
	sort.Ints(slice)
	mes("after sorted : ", slice)
}

func StringAdd()  {
	buffer := bytes.Buffer{}
	buffer.WriteString("aaa")
	buffer.WriteString("bbb")
	str := buffer.String()
	system.Println(str)
}

func StringTest() {
	str := "hello world"
	a := str[ len(str) / 2: ] + str[ : len(str) / 2 ]
	system.Println("result a : ", a)
}

func BufferUse()  {
	var buffer bytes.Buffer
	var buffer1 *bytes.Buffer = new(bytes.Buffer)
	buffer2 := bytes.NewBuffer([]byte{})
	system.Println("buffer : ", buffer)
	system.Println("buffer1 : ", buffer1)
	system.Println("buffer2 : ", buffer2)
}