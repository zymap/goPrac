package main

import (
	_ "./gohelloworld"
	"./io"
	"./learn"
	"./learn/collection"
	"./learn/gcandfinalize"
	"./learn/interface"
	lio "./learn/io"
	. "./learn/structandmethod"
	"container/ring"
	"fmt"
)

func myhelloworld() {
	Helloworld()
	helloworld.Helloworld()
	helloworld.Version()
}

func mylearn() {
	//learn.ValueTypeCopy()
	//learn.RefTypeCopy()
	//learn.VarTest()
	//learn.CreateRandom()
	learn.StringUtil()
}

func x() {
	i := 0
xxx:
	fmt.Println(i)
	i++
	if i < 20 {
		goto xxx
	}
}

func testFunction() {
	a := 10
	fmt.Println(&a)
	learn.FunctionX(&a)
	learn.FunctionY(a)
}

func OOP() {
	learn.MainFunc()
}

func array() {
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

const value = 100

func itoatest() {
	const (
		k1 = iota // 0
		k2        //iota 1
		k3        //iota 2
		k4 = value
		k5
		k6
		k7
		k8
	)
	const (
		b1 = iota //0
		b2        //iota 1
	)

	fmt.Println(k1, k2, k3, k4, k5, k6, k7, k8)
}

func iotest() {
	//io.Shell()
	io.IoOpt()
}

func slice() {
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

func reslice() {
	format("reslice 1")
	collection.Reslice1()
	format("reslice 2")
	collection.Reslice2()
	format("copy slice")
	collection.CopySlice()
}

func maptest() {
	format("map1")
	collection.Map1()
	format("map2")
	collection.Map2()
	format("map3")
	collection.Map3()
	format("mapslice")
	collection.MapSlice1()
	format("map4")
	collection.Map4()
}

func struct1() {
	format("struct1")
	Struct1()
	format("method1")
	Method1()
}

func other() {
	format("magic test")
	MagicTest()
	format("gc test")
	gcandfinalize.GCTest()
}

func interfacetest() {
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
	format("act")
	_interface.Act()
}

func miotest() {
	format("scanner1")
	lio.Scanner1()
	format("scanner2")
	lio.Scanner2()
}

func filetest() {
	format("input file1")
	lio.InputFile()
	format("input file2")
	lio.InputFile2()
	format("input file3")
	lio.InputFile3()
	format("input file4")
	lio.InputFile4()
}

func main() {
	other()

	r := ring.New(2)

}

func format(str string) {
	fmt.Print("====", str, "====")
	fmt.Print("====", str, "====\n")
}

// 483045022100ec1aa939322b57beac89e419356721a2e4ec0b96e2b11365cd260042ee43556c02205fc044071c68730ffb14fdbc7850b267a2fad1afd44d2f68efd7fa9c1171388b4141046d4da4a23215995d0a459070b0cbce06db3331bb66f7d94dfa98b137916a6c29d487a3a1f2b9d0acb4ad4a7dd56203f59731697ab4102110115d5e19482ffb24
// 483045022100ec1aa939322b57beac89e419356721a2e4ec0b96e2b11365cd260042ee43556c02205fc044071c68730ffb14fdbc7850b267a2fad1afd44d2f68efd7fa9c1171388b4141046d4da4a23215995d0a459070b0cbce06db3331bb66f7d94dfa98b137916a6c29d487a3a1f2b9d0acb4ad4a7dd56203f59731697ab4102110115d5e19482ffb24
