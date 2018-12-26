package hotupdate

import (
	"os"
	"log"
	"syscall"
)

func Start()  {
	if len(os.Args) == 1 {
		log.Fatal("input error")
	}

	method := os.Args[1]
	serverName := os.Args[2]

	if method == "reload" {
		fork, err := syscall.ForkExec(os.Args[0], os.Args)
	}

}