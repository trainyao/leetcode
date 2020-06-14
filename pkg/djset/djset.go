package djset

import (
	"errors"
)

type djset struct {
	parents   map[int]int
	setHeight map[int]int
	sets      map[int][]int
}

var AlreadyInOneSet = errors.New("already in one set")
var SetNotExists = errors.New("set not exists")

func NewDisJointSet() *djset {
	return &djset{
		parents:   make(map[int]int, 0),
		setHeight: make(map[int]int, 0),
		sets:      make(map[int][]int, 0),
	}
}

func (s *djset) FindSet(index int) (setId int, err error) {
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

func (s *djset) JoinSet(index int, relatedIndex int) (setId int) {
	relatedSet, errFindingRelatedIndexSet := s.FindSet(relatedIndex)
	// if relatedIndex not exist, create relatedIndex as a set first
	if errFindingRelatedIndexSet != nil {
		s.parents[relatedIndex] = relatedIndex
		s.setHeight[relatedIndex] = 1
		s.sets[relatedIndex] = []int{relatedIndex}
		errFindingRelatedIndexSet = nil
		relatedSet = relatedIndex
	}

	_, errFindingIndexSet := s.FindSet(index)

	// if both indices' set exist, it's a union operation
	if errFindingRelatedIndexSet == nil && errFindingIndexSet == nil {
		setId, _ = s.Union(index, relatedIndex)
		return
	}

	// set parent
	s.parents[index] = relatedSet
	s.sets[relatedSet] = append(s.sets[relatedSet], index)
	if s.setHeight[relatedSet] == 1 {
		s.setHeight[relatedSet]++
	}
	return relatedSet
}

func (s *djset) Union(index int, index2 int) (setId int, err error) {
	// check if this two index is in one set
	parent1, err := s.FindSet(index)
	if err != nil {
		return 0, err
	}
	parent2, err := s.FindSet(index2)
	if err != nil {
		return 0, err
	}

	if parent1 == parent2 {
		return parent2, AlreadyInOneSet
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

	// move 1's members into 2's
	for _, member := range s.sets[parent1] {
		s.sets[parent2] = append(s.sets[parent1], member)
	}
	// delete 1's set record
	delete(s.sets, parent1)

	return parent2, nil
}

func (s *djset) SetIds() (setIds []int) {
	for setId := range s.sets {
		setIds = append(setIds, setId)
	}
	return
}

func (s *djset) Members(setId int) []int {
	if members, ok := s.sets[setId]; ok {
		return members
	}

	return nil
}

func (s *djset) Sets() (setIdMappedMembers map[int][]int) {
	return s.sets
}
