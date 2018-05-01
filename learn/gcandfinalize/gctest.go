package gcandfinalize

import (
	"runtime"
	"fmt"
)

func GCTest()  {
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	fmt.Println("use memory : ", mem.Alloc / 1024, "kb")
}

func Finalize()  {
	t := new(int)
	*t = 1
	runtime.SetFinalizer(t,Finalize)
}