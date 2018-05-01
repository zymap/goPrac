package collection

import "fmt"

func printmap(mes string, m map[int]int)  {
	fmt.Println(mes, " is : ", m)
	fmt.Println(mes, " len : ", len(m))
	fmt.Println(mes, " cap : ")
}

func Map1()  {
	map1 := map[int]int{3:5, 4:5}
	map1[0] = 1
	map1[1] = 2
	printmap("map1", map1)

	//make map
	map2 := make(map[int]int)
	map22 := map2
	map2[1] = 2
	printmap("map2", map2)
	printmap("map22", map22)
	//update
	map22[2] = 20
	printmap("after update map2", map2)
	printmap("fater update map22", map22)

	map3 := make(map[int]int, 5)
	printmap("map3", map3)
}

func MapNewError()  {
	//map1 := new(map[int]int)
	//map1[0] = 1
}

func Map2()  {
	map1 := make(map[int][]int)
	slice := make([]int, 3)
	map1[0] = slice
	fmt.Println("map1 : ", map1)

	map2 := make(map[int]*[]int)
	slice1 := new([10]int)[:]
	slice1[0] = 1
	map2[0] = &slice1
	fmt.Println("map2 : ", map2)

	for k,v := range map2 {
		fmt.Println("map2 key : ", k, " ; ", "value : ", *v)
	}
}

func Map3()  {
	fmt.Println("func Map3 : key and value option include add, remove, exist...")
	map1 := map[int]int{1 : 1, 2 : 2, 3 : 3, 4 : 4}
	printmap("map1 : ", map1)
	//exist
	val1, isPresent := map1[0]
	fmt.Println("map1[0] exists : ", isPresent, " \t value : ", val1)
	val2, isExist := map1[1]
	fmt.Println("map1[1] exists : ", isExist, " \t value : ", val2)

	//remove
	delete(map1, 2)
	printmap("deleted map1[2]", map1)
}

func Map4()  {
	m := make(map[int]string)
	m[1] = "map4"
	fmt.Println("before update : ", m)
	mapsend(m)
	fmt.Println("update map : ", m)
}

func mapsend(m map[int]string)  {
	fmt.Println("enter mapsend : ", m)
	defer fmt.Println("exit mapsend : ", m)
	m[10] = "mapsend"
}

func MapSlice1()  {
	mapslice := make([]map[int]int,5)

	for k, _ := range mapslice {
		mapslice[k] = make(map[int]int, 2)
		mapslice[k][0] = 1
		mapslice[k][1] = 2
	}

	fmt.Println("mapslice : ", mapslice)
}