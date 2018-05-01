package learn

var a = 10

func VarTest() {
	n()
	mm()
	n()
}

func n()  {
	println(a)
}

func m()  {
	a := 20
	println(a)
}

func mm()  {
	a = 20
	println(a)
}