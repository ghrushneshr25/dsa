// @type: concept
// @index: 2
// @title: Doubly Linked List
// @tags: linked-list, fundamentals
// @difficulty: Easy

/*
@section: Description
A doubly linked list is a linear data structure where each node contains: Data, Pointer to the next node, Pointer to the previous node. The next of last node points to nil, indicating the end of the list. The previous of first node points towards nil.

@section: Structure
@section: Operations
*/

package linkedlist

import (
	"errors"
	"fmt"
)

// @structure Double List Node
// @description Represents a single node in the linked list.
type DoubleListNode struct {
	Data any
	Next *DoubleListNode
	Prev *DoubleListNode
}

// @structure DoublyLinkedList
// @description Represents the linked list with head pointer and size.
type DoublyLinkedList struct {
	Head *DoubleListNode
	Tail *DoubleListNode
	Size int
}

// @operation Traversal
// @description Traversing means visiting each node from head to nil.
// @time: O(n)
// @space: O(1)
func (list *DoublyLinkedList) Display() (err error) {
	if list.Head == nil {
		return errors.New("list is empty")
	}

	current := list.Head
	for current != nil {
		fmt.Printf("%+v <->", current.Data)
		current = current.Next
	}
	return
}

// @operation Increase Size
// @description Increase the size of linked list
func (list *DoublyLinkedList) IncreaseSize() {
	list.Size++
}

// @operation Decrease Size
// @description Decrease the size of linked list
func (list *DoublyLinkedList) DecreaseSize() {
	list.Size--
}

// @operation Get Size
// @description Get the size of linked list
func (list *DoublyLinkedList) GetSize() int {
	return list.Size
}

// @operation Insertion at Beginning
// @description Inserting the node the at the beginning of doubly linked list
// @time: O(1)
// @space: O(1)
func (list *DoublyLinkedList) InsertAtBeginning(data any) (err error) {
	newNode := &DoubleListNode{
		Data: data,
		Next: nil,
		Prev: nil,
	}

	defer list.IncreaseSize()

	if list.GetSize() == 0 {
		list.Head = newNode
		list.Tail = newNode
		return
	}

	list.Head.Prev = newNode
	newNode.Next = list.Head
	list.Head = newNode
	return
}

// @operation Insertion at End
// @description Inserting the node the at the end of doubly linked list
// @time: O(1)
// @space: O(1)
func (list *DoublyLinkedList) InsertAtEnd(data any) (err error) {
	newNode := &DoubleListNode{
		Data: data,
		Prev: nil,
		Next: nil,
	}

	defer list.IncreaseSize()

	if list.GetSize() == 0 {
		list.Head = newNode
		list.Tail = newNode
		return
	}

	list.Tail.Next = newNode
	newNode.Prev = list.Tail
	list.Tail = newNode

	return
}

// @operation Insert at Position
// @description Inserting the node at the given location in the linked list
// @time: O(n)
// @space: O(1)
func (list *DoublyLinkedList) Insert(position int, data any) (err error) {
	if position < 1 || position > list.GetSize()+1 {
		return fmt.Errorf("index out of bound: %d", position)
	}

	if position == 1 {
		return list.InsertAtBeginning(data)
	}

	if position == list.GetSize()+1 {
		return list.InsertAtEnd(data)

	}

	newNode := &DoubleListNode{
		Data: data,
		Prev: nil,
		Next: nil,
	}

	// optimization to bi directional insert
	current := list.Head
	if position <= list.GetSize()/2 {
		for i := 1; i < position; i++ {
			current = current.Next
		}
	} else {
		current = list.Tail
		for i := list.GetSize(); i >= position; i-- {
			current = current.Prev
		}
	}

	newNode.Prev = current.Prev
	newNode.Next = current

	current.Prev.Next = newNode
	current.Prev = newNode

	list.IncreaseSize()

	return
}

// @operation Deleting the first node
// @description Deleting the first node in the linked list
// @time: O(n)
// @space: O(1)
func (list *DoublyLinkedList) DeleteFirst() (data any, err error) {
	if list.Head == nil {
		return nil, fmt.Errorf("list is empty")
	}
	data = list.Head.Data
	if list.Head.Next == nil {
		list.Head = nil
		list.Tail = nil
	} else {
		list.Head = list.Head.Next
		list.Head.Prev = nil
	}
	list.DecreaseSize()
	return
}

// @operation Deleting the last node
// @description Deleting the last node in the linked list
// @time: O(n)
// @space: O(1)
func (list *DoublyLinkedList) DeleteLast() (data any, err error) {
	if list.Tail == nil {
		return nil, fmt.Errorf("list is empty")
	}
	data = list.Tail.Data
	if list.Tail.Prev == nil {
		list.Head = nil
		list.Tail = nil
	} else {
		list.Tail = list.Tail.Prev
		list.Tail.Next = nil
	}
	list.DecreaseSize()
	return
}

// @operation Deleting an intermediate node
// @description Deleting the node at a given location in the linked list
// @time: O(n)
// @space: O(1)
func (list *DoublyLinkedList) Delete(position int) (data any, err error) {
	if position < 1 || position > list.GetSize() {
		return nil, fmt.Errorf("index out of bound: %d", position)
	}

	// Case 1: Only one node
	if list.GetSize() == 1 {
		data := list.Head.Data
		list.Head = nil
		list.Tail = nil
		list.DecreaseSize()
		return data, nil
	}

	// Case 2: Delete head
	if position == 1 {
		data := list.Head.Data
		list.Head = list.Head.Next
		list.Head.Prev = nil
		list.DecreaseSize()
		return data, nil
	}

	// Case 3: Delete tail
	if position == list.GetSize() {
		data := list.Tail.Data
		list.Tail = list.Tail.Prev
		list.Tail.Next = nil
		list.DecreaseSize()
		return data, nil
	}

	current := list.Head

	if position <= list.GetSize()/2 {
		for i := 1; i < position; i++ {
			current = current.Next
		}
	} else {
		current = list.Tail
		for i := list.GetSize(); i > position; i-- {
			current = current.Prev
		}
	}

	data = current.Data

	current.Prev.Next = current.Next
	current.Next.Prev = current.Prev

	list.DecreaseSize()
	return
}
