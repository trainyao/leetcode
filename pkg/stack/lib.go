package stack

import "errors"

type Stack []interface{}

var StarkEmpty = errors.New("stack empty")

func (s *Stack) Push(i interface{}) {
	*s = append(*s, i)
}

func (s *Stack) Pop() (i interface{}, err error) {
	l := len(*s)
	if l == 0 {
		return 0, StarkEmpty
	}

	i = (*s)[l-1]
	*s = (*s)[:l-1]

	return i, nil
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}
