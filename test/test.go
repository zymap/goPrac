package main

type stu struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]*stu)
	stus := []stu{
		{"zhou", 24},
		{"li", 23},
		{"wang", 22},
	}

	for _, stu := range stus {
		m[stu.Name] = &stu
	}
}
