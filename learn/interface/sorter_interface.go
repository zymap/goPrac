package _interface

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(data Sorter)  {
	for i := 1; i < data.Len(); i++ {
		for j := 0; j < data.Len() - i; j++ {
			if data.Less(i + 1, i) {
				data.Swap(i, i + 1)
			}
		}
	}
}