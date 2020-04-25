package lib

import (
	"fmt"
	"testing"
)

func Test_stack(t *testing.T) {
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)

	i, err := s.Pop()
	if err != nil {
		t.Fatal(err)
	}
	j, err := s.Pop()
	if err != nil {
		t.Fatal(err)
	}
	k, err := s.Pop()
	if err != nil {
		t.Fatal(err)
	}

	r := fmt.Sprintf("%d%d%d", i, j, k)

	if r != "321" {
		t.FailNow()
	}

	_, err = s.Pop()
	if err == nil {
		t.Fatal("err should not be empty")
	}
}
