// @type: concept
// @index: 3
// @title: Circular Linked List
// @tags: linked-list, fundamentals
// @difficulty: Easy

/*
@section: Description
In a circular linked list, every node point to its next node in the sequence but the last node points to the first node in the list.

@section: Structure
@section: Operations
*/

package linkedlist

import (
	"errors"
	"fmt"
)

// @structure Circular List Node
// @description Represents a single node in the circular linked list.
type CircularListNode struct {
	Data any
	Next *CircularListNode
}

// @structure CircularLinkedList
// @description Represents the circular linked list with head pointer and size.
type CircularLinkedList struct {
	Size int
	Head *CircularListNode
}

// @operation Traversal
// @description Traversing means visiting each node of circular linked list.
// @time: O(n)
// @space: O(1)
func (list *CircularLinkedList) Display() (result string, err error) {
	if list.Head == nil {
		err = errors.New("list is empty")
		return
	}

	current := list.Head

	for {
		fmt.Printf("%+v", current.Data)
		current = current.Next
		if current == list.Head {
			break
		}
		fmt.Printf("->")
	}
	return
}

// @operation Increase Size
// @description Increase the size of linked list
func (list *CircularLinkedList) IncreaseSize() {
	list.Size++
}

// @operation Decrease Size
// @description Decrease the size of linked list
func (list *CircularLinkedList) DecreaseSize() {
	list.Size--
}

// @operation Get Size
// @description Get the size of linked list
func (list *CircularLinkedList) GetSize() int {
	return list.Size
}

// @operation Insert at Front
// @description Inserting the node the at the beginning of circular linked list
// @time: O(n)
// @space: O(1)
func (list *CircularLinkedList) InsertAtFront(data any) (err error) {
	newNode := &CircularListNode{
		Data: data,
	}
	newNode.Next = newNode

	if list.Head == nil {
		list.Head = newNode
		list.IncreaseSize()
		return
	}

	current := list.Head
	newNode.Next = current

	for current.Next != list.Head {
		current = current.Next
	}

	current.Next = newNode
	list.Head = newNode
	list.IncreaseSize()
	return
}

// @operation Insert at End
// @description Inserting the node the at the end of circular linked list
// @time: O(n)
// @space: O(1)
func (list *CircularLinkedList) InsertAtEnd(data any) (err error) {
	newNode := &CircularListNode{
		Data: data,
	}
	newNode.Next = newNode

	if list.Head == nil {
		list.Head = newNode
		list.IncreaseSize()
		return
	}

	current := list.Head

	for current.Next != list.Head {
		current = current.Next
	}
	newNode.Next = list.Head
	current.Next = newNode
	list.IncreaseSize()
	return
}

// @operation Insert at Position
// @description Inserting the node the at a given position of circular linked list
// @time: O(n)
// @space: O(1)
func (list *CircularLinkedList) InsertAtPosition(position int, data any) (err error) {
	newNode := &CircularListNode{
		Data: data,
	}
	newNode.Next = newNode

	if list.Head == nil {
		list.Head = newNode
		list.IncreaseSize()
		return
	}

	if position < 0 {
		err = fmt.Errorf("invalid position")
		return
	}
	position = (position-1)%list.GetSize() + 1
	if position == 1 {
		return list.InsertAtFront(data)
	}

	current := list.Head

	for i := 1; i < position-1; i++ {
		current = current.Next
	}

	newNode.Next = current.Next
	current.Next = newNode
	list.IncreaseSize()

	return
}
