package _interface

import (
	"fmt"
	"reflect"
)

func Reflect()  {
	x := 3.4
	fmt.Println(" x type : ", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("x vlue : ", v)
	fmt.Println("v type : ", v.Type())
	fmt.Println("v kind : ", v.Kind())
	fmt.Println("value : ", v.Float())
	fmt.Println("v.interface : ", v.Interface())

	fmt.Println("update : ")
	fmt.Println("can v be reset ? ", v.CanSet())
	v = reflect.ValueOf(&x)
	fmt.Println("&x; can v be reset ? ", v.CanSet())
	v = v.Elem()
	fmt.Println("v.elem(); can v be reset ? ", v.CanSet())
	v.SetFloat(3.1415926)
	fmt.Println("v reset : ", v)
}

type reflectStrcuct struct {
	id int
	name string
}

func (this reflectStrcuct) Show()  string {
	fmt.Println("id : ", this.id, "\tname : ", this.name)
	return "hello myshow"
}

var s interface{} = reflectStrcuct{1, "aaa"}

func Reflect1()  {
	test := s

	value := reflect.ValueOf(test)
	_type := reflect.TypeOf(test)
	value_type := value.Type()
	_kind := value.Kind()
	fmt.Println("test value : ", value)
	fmt.Println("test type : ", _type)
	fmt.Println("test value type : ", value_type)
	fmt.Println("test value kind : ", _kind)

	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("value filed %d : %v \n", i, value.Field(i))
	}

	result := value.Method(0).Call(nil)
	fmt.Println(result)
}

type T struct {
	A int
	B string
}

func Reflect2()  {
	t := T{1, "aaa"}
	s := reflect.ValueOf(&t).Elem()
	typeoft := s.Type()
	fmt.Println("type of t : ", typeoft)
	for i := 0; i < s.NumField(); i++ {
		fmt.Printf("%d : %s %s = %v\n", i, typeoft.Field(i).Name, s.Field(i).Type(), s.Field(i).Interface())
	}
	fmt.Printf("update\n")
	s.Field(0).SetInt(2)
	s.Field(1).SetString("bbb")
	fmt.Println("t : ", t)
}