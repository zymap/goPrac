
//web框架:beego

package learn

import (
	"fmt"
	//"time"
	"time"
)

/*
var a bool = true
var b bool = false
var c = 10
d := 20
c := 10

func basicone()  {
	e := 29


}
*/

func Basictype() {
	var i  = "abc"
	fmt.Println(i)

	x:= 1
	fmt.Println(x)

	var v string  = "abcdefg"
	fmt.Println(v)

	var inti int32 = 2
	fmt.Println(inti)
}

func Basiclanguage() {
	if 1== 1 {
		fmt.Println("this is true")
	}

}

func Compare(a int, b int) {
	if  a == b{
		fmt.Println(a, " is equals with ", b)
	} else if a < b {
		fmt.Println(a," is smaller than ", b)
	}else {
		fmt.Println(a," is bigger than ", b)
	}
}

func Showsomething()  {
	var a = [10][10] int{}
	for i:=0;i<10 ;i++  {
		for j:=0;j<10 ;j++  {
			if i == j {
				a[i][j] = 1
			}
			if i == 9 && j != 0 {
				a[i][j] = 1
			}

			if j == 0 && i != 0 {
				a[i][j] = 1
			}
		}
	}
	for i:=0;i<10 ;i++  {
		fmt.Println(a[i])
	}
}

func Mybreaktest(value int) {
	begin:for value < 20{
		fmt.Printf("this is the value: %d\n", value)
		if value == 10 {
			break
		}
		value--
		if value == 15 {
			fmt.Println("return to begin...")
			goto begin
		}
	}
	//goto begin
	fmt.Println("the function has been breaked...")
}

func Swap(a int,b int) (int, int)  {
	return b, a
}

func Max(a int,b int) int  {
	if a> b {
		return a
	}else {
		return b
	}
}

func Addsomething(s *int)  {
	fmt.Println(*s)
	var temp = *s
	for i:=0;i<10 ;i++  {
		temp +=temp
		fmt.Println(temp)
	}
	*s = temp
}

func Selectsomething(value *int) {

	switch *value {
	case 0:
		fmt.Println("this is ", value)
	case 1:
		fmt.Println("this is ", value)
	case 2:
		fmt.Println("this is ", value)
		fallthrough
	case 3:
		fmt.Println("this is ", value)
	default:
		fmt.Println("this is default")
	}
}



func bdtype()  {
	var i = 1
	print(i)
}

func learnmain() {
	bdtype()
}

func Arraypractice() {
	var a = []int{1,2, 3,4, 5, 6}
	fmt.Println(a)
	for i:=0; i < len(a) ;i++ {
		a[i] = a[i]+ 1
	}
	fmt.Println(a)
}

func Slice() {
	//var a = [] int {1,2,3,4,5, 6}
	//var b [] int = make(a, 1)
}

func Mystruct() {
	type Thing struct {
		string
	}

	type People struct {
		id int
		name string
		Thing
	}

	var p = People{1, "a",Thing{"aaa"}}
	fmt.Println(p)

	p.id = 2
	p.name = "b"
	fmt.Println(p)

	var c = struct {
		cid int
		cname string
	}{1,"child1"}

	fmt.Println(c)

	type Child struct {
		cid int
		cname string
	}

	var parent = struct {
		Child
		c Child
		id int
		name string
	}{}

	parent.cid = 1
	parent.cname= "cname"
	parent.c.cname = "ccname"
	parent.c.cid = 2
	parent.id = 3
	parent.name = "parentname"

	fmt.Println(parent)
}


func Mystructfunction() {
	var c = Car{1, "c1"}
	fmt.Println(c.GetId())
	fmt.Println(c.GetName())
	c.SetName("cc1")
	fmt.Println(c.GetName())
	c.Set_Name("ccc11")
	fmt.Println(c.GetName())

	var p = Product{Car{11,"11c"},1, "p111"}
	//var c = Car{}

	fmt.Println(p)
	fmt.Println(p.GetId())



}

type Car struct {
	id int
	name string
}

func (c Car) GetId() int {
	return c.id
}

func (c Car) GetName() string  {
	return c.name
}

func (c Car) SetName(name string){
	c.name = name
}

func (c *Car) Set_Name(name string){
	c.name = name
}

type Product struct {
	Car
	id int
	name string
}

func (p Product) GetId() int {
	return p.id
}


func Myexception() {
	defer func() {
		err:=recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	panic("error")
}


func My_exception(i int) {
	defer func() {
		err:=recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if i== 0 {
		panic("this is ")
	}

}

func T1(c *chan int)  {
	v:=<-*c
	fmt.Println(v)
	*c <- 1
}

func Myrouteine() {
	c := make(chan int)
	go T1(&c)
	c<-1;
	time.Sleep(3*time.Second)
	v,err:=<-c
	fmt.Println(v,err)
}

type P struct {
	s S
	id int
	name string
}

type S struct {
	f float32
}

func Myref() {
	p := P{s:S{1.1},id:1,name: "xxxx"}
	fmt.Println(p)
}

func Myselect() {
	mychan := make(chan int)
	go func() {
		mychan <- 1
	}()

	select {
	case j:=<-mychan:
		println(j)
	}
}