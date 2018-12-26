package main

import (
	"flag"
	"os"
)

func main() {
	s := flag.NewFlagSet("hhh", flag.ExitOnError)
	s.String("set1", "", "set 1")
	s.Parse(os.Args)

	s2 := flag.NewFlagSet("www", flag.ExitOnError)
	s2.String("set2", "", "set 2")
	s2.Parse(os.Args)
	flag.Parse()

}
