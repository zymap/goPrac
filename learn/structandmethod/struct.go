package structandmethod

import "fmt"

type People struct {
	id int
	name string
	live string
}

func printStruct(mes string, p People)  {
	fmt.Println(mes, " : ", p)
}

func Struct1()  {
	p1 := People{2, "jony", "china"}
	printStruct("p1", p1)

	p := new(People)
	p.id = 1
	p.name = "mike"
	p.live = "America"
	printStruct("p", *p)
}

type student struct {
	id int "this is student id"
}

func Struct2() {

}

