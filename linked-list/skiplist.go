// @type: concept
// @index: 6
// @title: Skip List
// @tags: linked-list, probabilistic, advanced
// @difficulty: Medium

/*
@section: Description
A Skip List is a probabilistic data structure that allows fast search, insertion, and deletion in O(log n) average time. It consists of multiple levels of linked lists, where higher levels allow skipping multiple nodes.

@section: Structure
@section: Operations
*/

package linkedlist

import (
	"fmt"
	"math/rand"
)

// @structure Skip List Node
type SkipListNode struct {
	Value   any
	Forward []*SkipListNode
}

// @structure SkipList
type SkipList struct {
	Head       *SkipListNode
	MaxLevel   int
	Level      int
	P          float64
	Comparator func(a, b any) int
}

// @operation Constructor
func NewSkipList(maxLevel int, p float64, cmp func(a, b any) int) *SkipList {
	head := &SkipListNode{
		Forward: make([]*SkipListNode, maxLevel),
	}

	return &SkipList{
		Head:       head,
		MaxLevel:   maxLevel,
		Level:      1,
		P:          p,
		Comparator: cmp,
	}
}

// @operation Random Level
func (list *SkipList) randomLevel() int {
	level := 1
	for rand.Float64() < list.P && level < list.MaxLevel {
		level++
	}
	return level
}

// @operation Search
func (list *SkipList) Search(target any) (bool, error) {
	if list.Comparator == nil {
		return false, fmt.Errorf("comparator not defined")
	}

	current := list.Head

	for i := list.Level - 1; i >= 0; i-- {
		for current.Forward[i] != nil &&
			list.Comparator(current.Forward[i].Value, target) < 0 {
			current = current.Forward[i]
		}
	}

	current = current.Forward[0]

	return current != nil &&
		list.Comparator(current.Value, target) == 0, nil
}

// @operation Insert
func (list *SkipList) Insert(value any) error {
	if list.Comparator == nil {
		return fmt.Errorf("comparator not defined")
	}

	update := make([]*SkipListNode, list.MaxLevel)
	current := list.Head

	// Step 1: Find insertion points
	for i := list.Level - 1; i >= 0; i-- {
		for current.Forward[i] != nil &&
			list.Comparator(current.Forward[i].Value, value) < 0 {
			current = current.Forward[i]
		}
		update[i] = current
	}

	// Step 2: Generate level
	newLevel := list.randomLevel()

	if newLevel > list.Level {
		for i := list.Level; i < newLevel; i++ {
			update[i] = list.Head
		}
		list.Level = newLevel
	}

	// Step 3: Create node
	newNode := &SkipListNode{
		Value:   value,
		Forward: make([]*SkipListNode, newLevel),
	}

	// Step 4: Insert
	for i := 0; i < newLevel; i++ {
		newNode.Forward[i] = update[i].Forward[i]
		update[i].Forward[i] = newNode
	}

	return nil
}

// @operation Delete
func (list *SkipList) Delete(value any) (bool, error) {
	if list.Comparator == nil {
		return false, fmt.Errorf("comparator not defined")
	}

	update := make([]*SkipListNode, list.MaxLevel)
	current := list.Head

	// Step 1: Find predecessors
	for i := list.Level - 1; i >= 0; i-- {
		for current.Forward[i] != nil &&
			list.Comparator(current.Forward[i].Value, value) < 0 {
			current = current.Forward[i]
		}
		update[i] = current
	}

	current = current.Forward[0]

	if current == nil ||
		list.Comparator(current.Value, value) != 0 {
		return false, nil
	}

	// Step 2: Remove node
	for i := 0; i < list.Level; i++ {
		if update[i].Forward[i] != current {
			break
		}
		update[i].Forward[i] = current.Forward[i]
	}

	// Step 3: Adjust level
	for list.Level > 1 &&
		list.Head.Forward[list.Level-1] == nil {
		list.Level--
	}

	return true, nil
}
