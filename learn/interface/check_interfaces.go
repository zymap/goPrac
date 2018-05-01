package _interface

import "fmt"

type ATest interface {
	A()
}

type BTest interface {
	B()
}

type CTest interface {
	ATest
}

type TestStruct struct {

}

func (this *TestStruct) A()  {
	fmt.Println("this is show TestStruct invoke A interface.")
}

func (this *TestStruct) B()  {
	fmt.Println("this is show TestStruct invoke B interface")
}

var Aintf ATest
var Bintf BTest
var Cintf CTest

func CheckI()  {
	a := new(TestStruct)
	b := new(TestStruct)

	Aintf = a
	Bintf = b

	if _type, IsExist := Aintf.(*TestStruct); IsExist {
		fmt.Println("IsExist : ", IsExist)
		fmt.Printf("the %v type is %T\n", _type, _type)
	}
}


//测试某个值是否实现了某个接口, 不成功打印不了
func Check2()  {
	Aintf = Cintf
	if _type, isExist := (Cintf).(ATest); isExist {
		fmt.Printf("the TestStruct implement : %T\n", _type)
	}
}