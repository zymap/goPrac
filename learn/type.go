package learn

//值类型和引用类型
//基本数据类型都是值类型 在进行赋值的时候是通过内存的拷贝完成的
//复杂的数据需要多个字 通过copy引用地址

import (
	system "fmt"
)

func ValueTypeCopy()  {
	println("this is value type copy...")
	a := 1
	system.Println("this is before copy ",a)
	aa := a
	system.Println("this is aa: ", aa)
	a = 10
	system.Println("update a : ", a)
	system.Println("aa : ", aa)
}

func RefTypeCopy()  {
	a := 10
	ra := &a
	system.Println("this is ref ra: ",*ra)
	rb := ra
	system.Println("this is ref rb: ",*rb)
	a = 20
	system.Println("this is after update ra: ",*ra)
	system.Println("this is after update rb: ",*rb)
}