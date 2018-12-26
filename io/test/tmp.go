package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	name, err := ioutil.TempDir("./", "copernicus")
	defer os.Remove(name)
	fmt.Println(name, err)

	time.Sleep(time.Minute)

}
