// @type: concept
// @index: 5
// @title: Unrolled Linked List
// @tags: linked-list, advanced, memory-optimization, cache-friendly
// @difficulty: Medium

/*
@section: Description
An Unrolled Linked List is a variation of a linked list where each node stores multiple elements in a contiguous block (array) instead of a single element. This reduces pointer overhead and improves cache locality, making traversal and sequential access more efficient than traditional linked lists.
*/

package linkedlist

import (
	"errors"
	"fmt"
)

// @structure Unrolled List Node
// @description Represents a node in unrolled linked list storing multiple elements.
type UnrolledListNode struct {
	Elements []any             // underlying storage (capacity fixed per node)
	Count    int               // number of elements currently used
	Next     *UnrolledListNode // pointer to next node
}

// @structure UnrolledLinkedList
// @description Represents the unrolled linked list with head pointer, capacity, and size.
type UnrolledLinkedList struct {
	Head       *UnrolledListNode // first node
	Capacity   int               // max elements per node
	Size       int               // total elements in list
	Comparator func(a, b any) int
}

// @operation Create Node
// @description Creates a new node with predefined capacity
func (list *UnrolledLinkedList) NewNode() *UnrolledListNode {
	return &UnrolledListNode{
		Elements: make([]any, list.Capacity),
		Count:    0,
		Next:     nil,
	}
}

// @operation Traversal
// @description Traversing means visiting each element in the unrolled linked list.
// Each node contains multiple elements, so traversal involves iterating through
// each node and then iterating through its stored elements sequentially.
// @time: O(n)
// @space: O(1)
func (list *UnrolledLinkedList) Display() (result string, err error) {
	if list.Head == nil {
		err = errors.New("empty list")
		return
	}
	current := list.Head
	for current != nil {

		for i := 0; i < current.Count; i++ {
			result = fmt.Sprintf("%+v -> %+v", result, current.Elements[i])
		}
		current = current.Next
	}
	return
}

// @operation Search
// @description Searches for a target element in the unrolled linked list.
// Uses block-level pruning by comparing with the maximum element of each block.
// Assumes elements inside each node are sorted.
// @time: O(√n)
// @space: O(1)

func (list *UnrolledLinkedList) Search(target any) (found bool, err error) {
	if list.Head == nil {
		return false, errors.New("list is empty")
	}

	if list.Comparator == nil {
		return false, fmt.Errorf("comparator not defined")
	}

	current := list.Head

	for current != nil {
		if current.Count == 0 {
			current = current.Next
			continue
		}

		blockMax := current.Elements[current.Count-1]

		if list.Comparator(target, blockMax) <= 0 {
			break
		}
		current = current.Next
	}

	if current == nil {
		return false, nil
	}

	for i := 0; i < current.Count; i++ {
		if list.Comparator(current.Elements[i], target) == 0 {
			return true, nil
		}
	}

	return false, nil
}

// @operation Insert
// @description Inserts an element while maintaining sorted order across blocks.
// If the node is full, it splits the node and redistributes elements.
// @time: O(n/k + k)
// @space: O(1)
func (list *UnrolledLinkedList) Insert(data any) (err error) {
	if list.Comparator == nil {
		return fmt.Errorf("comparator not defined")
	}

	// Case 1: Empty list
	if list.Head == nil {
		node := list.NewNode()
		node.Elements[0] = data
		node.Count = 1
		list.Head = node
		list.Size++
		return nil
	}

	current := list.Head

	for current != nil {
		if current.Count == 0 {
			break
		}
		blockMax := current.Elements[current.Count-1]
		if list.Comparator(data, blockMax) <= 0 {
			break
		}
		if current.Next == nil {
			break
		}
		current = current.Next
	}

	pos := 0
	for pos < current.Count && list.Comparator(current.Elements[pos], data) < 0 {
		pos++
	}

	if current.Count < list.Capacity {
		for i := current.Count; i > pos; i-- {
			current.Elements[i] = current.Elements[i-1]
		}
		current.Elements[pos] = data
		current.Count++
		list.Size++
		return nil
	}

	mid := list.Capacity / 2
	newNode := list.NewNode()

	j := 0
	for i := mid; i < current.Count; i++ {
		newNode.Elements[j] = current.Elements[i]
		j++
	}

	newNode.Count = current.Count - mid
	current.Count = mid

	newNode.Next = current.Next
	current.Next = newNode

	if list.Comparator(data, current.Elements[current.Count-1]) > 0 {
		current = newNode
		pos = 0
		for pos < current.Count && list.Comparator(current.Elements[pos], data) < 0 {
			pos++
		}
	} else {
		pos = 0
		for pos < current.Count && list.Comparator(current.Elements[pos], data) < 0 {
			pos++
		}
	}

	for i := current.Count; i > pos; i-- {
		current.Elements[i] = current.Elements[i-1]
	}
	current.Elements[pos] = data
	current.Count++
	list.Size++
	return nil
}

// @operation Delete (Sorted + Rebalance)
// @description Deletes an element while maintaining sorted order.
// Rebalances nodes using borrow or merge if underflow occurs.
// @time: O(n/k + k)
// @space: O(1)
func (list *UnrolledLinkedList) Delete(target any) (success bool, err error) {
	if list.Head == nil {
		return false, fmt.Errorf("list is empty")
	}

	if list.Comparator == nil {
		return false, fmt.Errorf("comparator not defined")
	}

	var prev *UnrolledListNode
	current := list.Head

	// Step 1: Find correct block
	for current != nil {
		if current.Count == 0 {
			prev = current
			current = current.Next
			continue
		}

		blockMax := current.Elements[current.Count-1]
		if list.Comparator(target, blockMax) <= 0 {
			break
		}

		prev = current
		current = current.Next
	}

	if current == nil {
		return false, nil
	}

	// Step 2: Find element inside node
	idx := -1
	for i := 0; i < current.Count; i++ {
		if list.Comparator(current.Elements[i], target) == 0 {
			idx = i
			break
		}
	}

	if idx == -1 {
		return false, nil
	}

	// Step 3: Delete + shift left
	for i := idx; i < current.Count-1; i++ {
		current.Elements[i] = current.Elements[i+1]
	}
	current.Count--
	list.Size--

	// Step 4: If node is still sufficiently full → done
	minThreshold := list.Capacity / 2
	if current.Count >= minThreshold || current.Next == nil {
		// Special case: empty head
		if current.Count == 0 && prev == nil {
			list.Head = current.Next
		}
		return true, nil
	}

	next := current.Next

	// Step 5: Try borrow from next
	if next.Count > minThreshold {
		// move first element of next → current
		current.Elements[current.Count] = next.Elements[0]
		current.Count++

		// shift next left
		for i := 0; i < next.Count-1; i++ {
			next.Elements[i] = next.Elements[i+1]
		}
		next.Count--

		return true, nil
	}

	// Step 6: Merge with next
	for i := 0; i < next.Count; i++ {
		current.Elements[current.Count+i] = next.Elements[i]
	}
	current.Count += next.Count
	current.Next = next.Next

	return
}
