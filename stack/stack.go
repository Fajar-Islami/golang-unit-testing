package stack

type Stack struct {
	arr []int
}

func New() *Stack {
	return &Stack{
		arr: make([]int, 0),
	}
}

func (s Stack) IsEmpty() bool {
	return len(s.arr) == 0
}

func (s Stack) Size() int {
	return len(s.arr)
}
