package structandmethod

import "fmt"

type T struct {
	id int
	name string
}

func (this *T) setId(id int)  {
	this.id = id
}

func (this *T) setName(name string)  {
	this.name = name
}

func (this *T) getName() (string) {
	return this.name
}

func (this *T) getId() int  {
	return this.id
}

func Method1()  {
	t := new(T)
	t.setId(1)
	t.setName("aaa")
	fmt.Println(t.getId(), "\t", t.getName())
	//fmt.Println(t.getName())

	tt := T{}
	tt.setId(2)
	tt.setName("bbb")
	fmt.Println(tt.getId(), "\t", tt.getName())
}