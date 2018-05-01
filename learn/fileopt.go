package learn

import (
	"os"
	"fmt"
)

var base_url = "/home/zy/myuser/go/goproject/goPrac/file"
var filename = "/output.txt"
var writefile = base_url + filename

func Writefile() {


	fout,err := os.Create(writefile)
	defer fout.Close()

	if err!= nil {
		fmt.Println(writefile, err)
		return
	}

	fout.WriteString("this is my output text...\n")
	fout.Write([]byte("byte write\n"))
}

func Readfile()  {
	fin,err := os.Open(writefile)
	defer fin.Close()
	if err != nil {
		 fmt.Println(writefile, err)
		 return
	}

	buffer := make([]byte, 1024)

	for{
		n,_ := fin.Read(buffer)
		if 0== n {
			break
		}
		os.Stdout.Write(buffer[:n])
	}
}

func File_opt()  {
	Readfile()
	//Writefile()
}