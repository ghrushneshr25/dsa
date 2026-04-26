// @type: concept
// @title: Singly Linked List
// @tags: linked-list, fundamentals
// @difficulty: Easy

/*
@section: Description
A singly linked list is a linear data structure where each node contains:
- Data
- Pointer to the next node

The last node points to nil, indicating the end of the list.
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
func (list *LinkedList) Display() error {
	if list.Head == nil {
		return errors.New("list is empty")
	}

	current := list.Head

	for current != nil {
		fmt.Printf("%+v -> ", current.Data)
		current = current.Next
	}
	return nil
}

// @operation Insertion at Beginning
// @description Inserting the node the at the beginning of linked list
// @time: O(1)
// @space: O(1)
func (list *LinkedList) InsertAtBeginning(data any) error {
	node := &ListNode{
		Data: data,
		Next: nil,
	}
	if list.Head == nil {
		list.Head = node
	} else {
		node.Next = list.Head
		list.Head = node
	}
	list.Size++
	return nil
}
