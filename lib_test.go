package main

import (
	"fmt"
	"testing"
)

func Test_stack(t *testing.T) {
	s := stack{}
	s.push(1)
	s.push(2)
	s.push(3)

	i, err := s.pop()
	if err != nil {
		t.Fatal(err)
	}
	j, err := s.pop()
	if err != nil {
		t.Fatal(err)
	}
	k, err := s.pop()
	if err != nil {
		t.Fatal(err)
	}

	r := fmt.Sprintf("%d%d%d", i, j, k)

	if r != "321" {
		t.FailNow()
	}

	_, err = s.pop()
	if err == nil {
		t.Fatal("err should not be empty")
	}
}
