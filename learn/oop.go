package learn

import "fmt"

type Student struct{
	id int
	name string
}

func (this Student)SetId(id int)  {
	this.id = id
}

func (this *Student)setId(id int)  {
	this.id = id
}

func (this Student)SetName(name string)  {
	this.name = name
}

func (this *Student)setName(name string)  {
	this.name = name
}

func (this Student)getId() (id int)  {
	id = this.id
	return
}

func (this Student)getName() (name string)  {
	name = this.name
	return
}

func StudentTest()  {
	stu := Student{}
	stu.setId(1)
	stu.setName("abc")
	fmt.Println(stu)
	fmt.Println(stu.getId(),"\t",stu.getName())
	fmt.Println("=====split=====")
	sstu := Student{}
	(*Student).setName(&sstu,"aabbcc")
	(*Student).setId(&sstu, 1000)
	fmt.Println(sstu)
	fmt.Println(sstu.getName(),"\t",sstu.getId())
}