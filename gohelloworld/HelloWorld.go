package gohelloworld

import (
	system "fmt"
	"runtime"
)

func Helloworld()  {
	system.Println("hello world")
}

func Version() {
	system.Println("go verison: ",runtime.Version())
}