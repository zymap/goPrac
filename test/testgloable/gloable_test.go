package testgloable

import (
	"fmt"
	"testing"
)

func initArgs() *Gloable {
	return GetInstance()
}

func TestGetInstance(t *testing.T) {
	g := initArgs()
	fmt.Println(g)
}

func TestGetInstance2(t *testing.T) {
	g := initArgs()
	fmt.Println(g)
}
