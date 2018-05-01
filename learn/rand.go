package learn

import (
	"math/rand"
	system "fmt"
	"time"
)

func CreateRandom()  {
	for i := 0; i < 10; i++  {
		a := rand.Int()
		system.Print(a,"\t")
	}
	system.Println()
	for i := 0; i < 5; i++  {
		a := rand.Intn(8)
		system.Print(a,"\t")
	}
	system.Println()

	seed := int64(time.Now().Nanosecond())
	rand.Seed(seed)
	for i := 0; i < 10; i++ {
		system.Print(rand.Intn(100),"\t")
	}
	system.Println()
}