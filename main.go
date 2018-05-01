package main

import (
	helloworld "./gohelloworld"
	"./learn"
	"./learn/collection"
	"fmt"
	"./io"
	. "./learn/structandmethod"
	"./learn/gcandfinalize"
	"./learn/interface"
)

func myhelloworld()  {
	helloworld.Helloworld()
	helloworld.Version()
}

func mylearn()  {
	//learn.ValueTypeCopy()
	//learn.RefTypeCopy()
	//learn.VarTest()
	//learn.CreateRandom()
	learn.StringUtil()
}

func x()  {
	i := 0
xxx:
	fmt.Println(i)
	i++;
	if i < 20 {
		goto xxx
	}
}

func testFunction()  {
	a := 10
	fmt.Println(&a)
	learn.FunctionX(&a)
	learn.FunctionY(a)
}

func OOP()  {
	learn.MainFunc()
}

func array()  {
	collection.Array1()
	fmt.Println("===array11===")
	collection.Array11()
	fmt.Println("===array3===")
	//learn.Array3()
	fmt.Println("====array4====")
	collection.Array4()
	fmt.Println("====array5====")
	collection.Array5()
}

const value  = 100

func itoatest()  {
	const (
		k1 = iota  // 0
		k2			//iota 1
		k3			//iota 2
		k4 = value
		k5
		k6
		k7
		k8
	)
	const(
		b1 = iota  //0
		b2  		//iota 1
	)

	fmt.Println(k1,k2,k3,k4,k5,k6,k7,k8)
}

func iotest()  {
	//io.Shell()
	io.IoOpt()
}

func slice()  {
	fmt.Println("====slice 1========slice 1====")
	collection.Slice1()
	fmt.Println("====slice 2========slice 2====")
	collection.Slice2()
	fmt.Println("====slice 3========slice 3====")
	collection.Slice3()
	fmt.Println("====slice 4========slice 4====")
	collection.Slice4()
	fmt.Println("====slice 5========slice 5====")
	collection.Slice5()
	fmt.Println("====string====")
	collection.StringAdd()
	format("string test")
	collection.StringTest()
	fmt.Println("====bytes use====")
	collection.BufferUse()
	format("sort slice")
	collection.SortSlice()
}

func reslice()  {
	format("reslice 1")
	collection.Reslice1()
	format("reslice 2")
	collection.Reslice2()
	format("copy slice")
	collection.CopySlice()
}

func maptest()  {
	format("map1")
	collection.Map1()
	format("map2")
	collection.Map2()
	format("map3")
	collection.Map3()
	format("mapslice")
	collection.MapSlice1()
}

func struct1()  {
	format("struct1")
	Struct1()
	format("method1")
	Method1()
}

func other()  {
	format("magic test")
	MagicTest()
	format("gc test")
	gcandfinalize.GCTest()
}

func interfacetest()  {
	format("type interface")
	_interface.TypeInterface()
	format("type interface 1")
	_interface.TypeInterface1()
	format("checkI")
	_interface.CheckI()
	format("check2")
	_interface.Check2()
	format("method set")
	_interface.MethodSet()
	format("myvector")
	_interface.MyVector()
	format("copy to nil interface slice")
	_interface.CopyToNilInterface()
	format("reflect")
	_interface.Reflect()
	format("reflect1")
	_interface.Reflect1()
	format("reflect2")
	_interface.Reflect2()
}

func main() {
	interfacetest()
}

func format(str string)  {
	fmt.Print("====", str, "====")
	fmt.Print("====", str, "====\n")
}
