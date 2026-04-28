// @type: concept
// @title: Singly Linked List
// @tags: linked-list, fundamentals
// @difficulty: Easy

/*
@section: Description
A singly linked list is a linear data structure where each node contains: Data and Pointer to the next node. The last node points to nil, indicating the end of the list.

@section: Structure
@section: Operations
*/

package linkedlist

import (
	"errors"
	"fmt"
)

// @structure List Node
// @description Represents a single node in the linked list.
type ListNode struct {
	Data any
	Next *ListNode
}

// @structure Linked List
// @description Represents the linked list with head pointer and size.
type LinkedList struct {
	Head *ListNode
	Size int
}

// @operation Traversal
// @description Traversing means visiting each node from head to nil.
// @time: O(n)
// @space: O(1)
func (list *LinkedList) Display() (err error) {
	if list.Head == nil {
		return errors.New("list is empty")
	}

	current := list.Head

	for current != nil {
		fmt.Printf("%+v -> ", current.Data)
		current = current.Next
	}
	return
}

// @operation Increase Size
// @description Increase the size of linked list
func (list *LinkedList) IncreaseSize() {
	list.Size++
}

// @operation Decrease Size
// @description Decrease the size of linked list
func (list *LinkedList) DecreaseSize() {
	list.Size--
}

// @operation Get Size
// @description Get the size of linked list
func (list *LinkedList) GetSize() int {
	return list.Size
}

// @operation Insertion at Beginning
// @description Inserting the node the at the beginning of linked list
// @time: O(1)
// @space: O(1)
func (list *LinkedList) InsertAtBeginning(data any) (err error) {
	node := &ListNode{
		Data: data,
		Next: nil,
	}

	defer list.IncreaseSize()

	if list.Head == nil {
		list.Head = node
	} else {
		node.Next = list.Head
		list.Head = node
	}
	return
}

// @operation Insertion at End
// @description Inserting the node the at the end of linked list
// @time: O(1)
// @space: O(1)
func (list *LinkedList) InsertAtEnd(data any) (err error) {
	node := &ListNode{
		Data: data,
		Next: nil,
	}

	defer list.IncreaseSize()

	if list.Head == nil {
		list.Head = node
	} else {
		current := list.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
	return
}

// @operation Insert at Position
// @description Inserting the node at the given location in the linked list
// @time: O(n)
// @space: O(1)
func (list *LinkedList) Insert(position int, data any) (err error) {
	if position < 1 || position > list.Size+1 {
		return fmt.Errorf("index out of bounds: %d", position)
	}

	node := &ListNode{
		Data: data,
		Next: nil,
	}

	defer list.IncreaseSize()

	if position == 1 {
		node.Next = list.Head
		list.Head = node

		return nil
	}

	var current *ListNode
	current = list.Head

	for i := 1; i < position-1; i++ {
		if current == nil {
			return errors.New("unexpected nil during traversal")
		}
		current = current.Next
	}

	node.Next = current.Next
	current.Next = node
	return
}

// @operation Deleting the first node
// @description Deleting the first node in the linked list
// @time: O(n)
// @space: O(1)
func (list *LinkedList) DeleteFirst() (data any, err error) {
	if list.Head == nil {
		return nil, errors.New("empty linked list")
	}

	defer list.DecreaseSize()

	data = list.Head.Data
	list.Head = list.Head.Next
	return
}

// @operation Deleting the last node
// @description Deleting the last node in the linked list
// @time: O(n)
// @space: O(1)
func (list *LinkedList) DeleteLast() (data any, err error) {
	if list.Head == nil {
		return nil, errors.New("empty linked list")
	}

	defer list.DecreaseSize()

	if list.Head.Next == nil {
		data = list.Head.Data
		list.Head = nil
		return
	}

	current := list.Head
	for current.Next.Next != nil {
		current = current.Next
	}

	data = current.Next.Data
	current.Next = nil
	return
}

// @operation Deleting an intermediate node
// @description Deleting the node at a given location in the linked list
// @time: O(n)
// @space: O(1)
func (list *LinkedList) Delete(position int) (data any, err error) {
	if position < 1 || position > list.GetSize() {
		return nil, fmt.Errorf("index out of bound: %d", position)
	}

	defer list.DecreaseSize()

	if position == 1 {
		data = list.Head.Data
		list.Head = list.Head.Next
		return
	}

	current := list.Head
	for i := 1; i < position-1; i++ {
		current = current.Next
	}

	if current.Next == nil {
		return nil, fmt.Errorf("unexpected nil node at position %d", position)
	}

	data = current.Next.Data
	current.Next = current.Next.Next
	return
}
