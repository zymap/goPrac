package someclass

type test struct {
	id int
}

func New() *test {
	return &test{1}
}

func (t *test) set(id int) {
	t.id = id
}

func (t *test) getId() int {
	return t.id
}

func (t *test) GetAll() int {
	return t.id
}
