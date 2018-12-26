package main

import (
	"os"
	"time"
)

func main() {
	os.Open("/Users/zhangyong/GoProjects/goPrac/test/aaa/a.tmp")
	time.Sleep(time.Second * 10)
	defer os.Remove("/Users/zhangyong/GoProjects/goPrac/test/aaa/a.tmp")
}
