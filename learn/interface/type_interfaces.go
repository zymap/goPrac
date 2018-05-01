package _interface

import (
	"math"
	"fmt"
)

type shaper interface {
	Area() float32
}

type square struct {
	side float32
}

type circle struct {
	radius float32
}

func (this *square) Area() float32  {
	return this.side * this.side
}

func (this *circle) Area() float32  {
	return this.radius * this.radius * math.Pi
}

func TypeInterface()  {
	var areaIntf shaper
	sq1 := new(square)
	sq1.side = 5

	areaIntf = sq1
	if t, ok := areaIntf.(*square); ok {
		fmt.Printf("the type of areaintf is : %T\n", t)
	} else {
		fmt.Println("areaintf does not contain a variable of type circle")
	}
	//fmt.Printf("the areainterf's type is : %T\n", areaIntf.(type))
	switch t := areaIntf.(type) {
	case *square:
		fmt.Printf("type square %T with value %v\n", t, t)
	case *circle:
		fmt.Printf("type circle %T with value %v \n", t, t)
	case nil:
		fmt.Printf("nil value : nothing to check ?\n")
	default:
		fmt.Printf("unexpected type %T \n", t)
	}
}

func TypeInterface1()  {
	c := circle{2}
	s := square{1}

	shapes := []shaper{&c, &s}
	for _, v := range shapes {
		fmt.Printf("the shape is : %T\n", v)
		fmt.Println("the shape area is : ", v.Area())
	}
}