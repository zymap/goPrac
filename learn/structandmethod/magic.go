package structandmethod

import "fmt"

type Base struct {
}

func (Base) Magic()  {
	fmt.Println("base magic")
}

func (this Base) MoreMagic()  {
	this.Magic()
	this.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic()  {
	fmt.Println("voodoo magic")
}

func MagicTest()  {
	v := new(Voodoo)
	v.Magic()
	//v.MoreMagic()
}