package djset

import (
	"errors"
)

type Djset struct {
	parents   map[int]int
	setHeight map[int]int
}

var AlreadyInOneSet = errors.New("already in one set")
var SetNotExists = errors.New("set not exists")

func NewDisJointSet() *Djset {
	return &Djset{
		parents:   make(map[int]int, 0),
		setHeight: make(map[int]int, 0),
	}
}

func (s *Djset) FindSet(index int) (setId int, err error) {
	// not exist, index is a new set, setId is itself
	if _, ok := s.parents[index]; !ok {
		return -1, SetNotExists
	}

	for {
		// find index's parent is itself
		if s.parents[index] == index {
			return index, nil
		}

		index = s.parents[index]
	}
}

func (s *Djset) JoinSet(index int, relatedIndex int) {
	relatedSet, errFindingRelatedIndexSet := s.FindSet(relatedIndex)
	// if relatedIndex not exist, create relatedIndex as a set first
	if errFindingRelatedIndexSet != nil {
		s.parents[relatedIndex] = relatedIndex
		s.setHeight[relatedIndex] = 1
		errFindingRelatedIndexSet = nil
		relatedSet = relatedIndex
	}

	_, errFindingIndexSet := s.FindSet(index)

	// if both indices' set exist, it's a union operation
	if errFindingRelatedIndexSet == nil && errFindingIndexSet == nil {
		_ = s.Union(index, relatedIndex)
		return
	}

	// set parent
	s.parents[index] = relatedSet
	if s.setHeight[relatedSet] == 1 {
		s.setHeight[relatedSet]++
	}
}

func (s *Djset) Union(index int, index2 int) error {
	// check if this two index is in one set
	parent1, err := s.FindSet(index)
	if err != nil {
		return err
	}
	parent2, err := s.FindSet(index2)
	if err != nil {
		return err
	}

	if parent1 == parent2 {
		return AlreadyInOneSet
	}

	// 1's is higher, use 1 as 2's parent
	if s.setHeight[parent1] > s.setHeight[parent2] {
		// by switching 2 and 1
		parent1, parent2 = parent2, parent1
	}

	// 2's height >= 1, use 2 as 1's parent
	s.parents[parent1] = parent2

	// if height is the same, parent's height++
	if s.setHeight[parent1] == s.setHeight[parent2] {
		s.setHeight[parent2]++
	}

	// delete 1's height record
	delete(s.setHeight, parent1)

	return nil
}
