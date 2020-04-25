package main

import "fmt"

type stack []interface{}

func (s *stack) push(i interface{}) {
	*s = append(*s, i)
}

func (s *stack) pop() (i interface{}, err error) {
	l := len(*s)
	if l == 0 {
		return 0, fmt.Errorf("stack empty")
	}

	i = (*s)[l-1]
	*s = (*s)[:l-1]

	return i, nil
}

func (s *stack) empty() bool {
	return len(*s) == 0
}
