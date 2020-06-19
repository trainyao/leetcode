package kmp

import "errors"

type KMPTable struct {
	next    []int
	cursorI int
	cursorJ int
	s       string

	// for break point building
	cursorIsMatching bool
	matchingCounter  int
}

func BuildKMP(s string, buildLen ...int) (*KMPTable, error) {
	t := &KMPTable{
		next: []int{},
		// starting at 1 as j
		cursorI: 0,
		cursorJ: 1,
		s:       s,
	}
	lenToBuild := len(s)
	if len(buildLen) == 1 {
		lenToBuild = buildLen[0]
	}

	// next[0] = -1
	// next[1] = 0
	t.next = append(t.next, -1)
	t.next = append(t.next, 0)

	err := t.build(lenToBuild)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (t *KMPTable) build(buildLen int) (err error) {
	if t.cursorJ >= len(t.s) {
		return errors.New("already reach string maxlen")
	}

	for t.cursorJ < buildLen {
		// firstly mismatch
		for !t.cursorIsMatching && t.s[t.cursorI] != t.s[t.cursorJ] {
			// if is about to reach max len
			if t.cursorJ+1 >= buildLen {
				return
			}

			t.cursorJ++
			t.next = append(t.next, 0)
		}

		// found match position
		// when match, start count match len
		t.cursorIsMatching = true
		for t.cursorIsMatching && t.s[t.cursorI] == t.s[t.cursorJ] {
			// if is about to reach max len
			if t.cursorJ+1 >= buildLen {
				return
			}
			t.cursorI++
			t.cursorJ++

			t.matchingCounter++
			t.next = append(t.next, t.matchingCounter)
		}
		// i, j = next mismatch index, continue to loop, till j = lenToBuild
		t.matchingCounter = 0
		t.cursorIsMatching = false
	}

	return
}

func (t *KMPTable) ContinueBuild(buildLen ...int) error {
	lenToBuild := len(t.s)
	if len(buildLen) == 1 {
		lenToBuild = buildLen[0]
	}
	return t.build(lenToBuild)
}
