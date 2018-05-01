package database

import (
	"os"
	"strings"
	"strconv"
	"fmt"
)

var fileurl = "./file/output.txt"

func Read() {
	fin,err := os.Open(fileurl)
	defer fin.Close()
	if err != nil {
		println(fileurl,err)
		return
	}

	line := ""
	buf := make([]byte, 1)
	for{
		n,_ := fin.Read(buf)
		if 0== n {
			break
		}
		if string(buf[:n]) == "\n" {
			str_util(line)
			line = ""
			continue
		}
		line = line + string(buf[:n])
	}
}

func str_util(str string) {
	println("str_util")
	println(str)
	stu := str_to_jdbcobj(str)
	//if stu !=  {
	//
	//}
	fmt.Println(stu)
}

func produce(stu Student) {

}

func consumer() {
	//databaseopt
}

func str_to_jdbcobj(str string) Student {
	strs :=strings.Split(str, " ")
	id,err := strconv.Atoi(strs[0])
	if err != nil {
		println(err)
		return Student{}
	}
	return Student{id,strs[1]}
}