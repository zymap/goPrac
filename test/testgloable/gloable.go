package testgloable

type Gloable struct {
	Flag bool
	Name string
}

var gloable *Gloable

func GetInstance() *Gloable {
	if gloable == nil {
		gloable = &Gloable{
			false,
			"hello",
		}
		return gloable
	}
	return initInstance()
}

func initInstance() *Gloable {
	gloable = &Gloable{
		true,
		"hello",
	}
	return gloable
}
