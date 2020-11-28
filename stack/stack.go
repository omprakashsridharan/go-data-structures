package stack

type Stack []int

func New() *Stack {
	var newStack Stack
	return &newStack
}

func (s *Stack) Height() int {
	return len(*s)
}

func (s *Stack) Peek() int {
	h := s.Height()
	if h > 0 {
		return (*s)[h-1]
	} else {
		return -1
	}
}

func (s *Stack) Push(data int) {
	*s = append(*s, data)
}

func (s *Stack) Pop() (int, bool) {
	h := s.Height()
	if h == 0 {
		return -1, false
	} else {
		topIndex := h - 1
		topItem := (*s)[topIndex]
		*s = (*s)[:topIndex]
		return topItem, true
	}
}

func (s *Stack) Print() {
	println("Printing Stack")
	h := s.Height()
	for i := h - 1; i >= 0; i-- {
		println("\t------\t")
		println("\t|", (*s)[i], "|\t")
		println("\t------\t")
	}
}
