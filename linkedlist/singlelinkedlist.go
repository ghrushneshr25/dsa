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

@subsection: ListNode
Represents a single node in the linked list.

@subsection: LinkedList
Represents the linked list with head pointer and size.

@section: Operations

@subsection: Traversal
Traversing means visiting each node from head to nil.

@subsection: Insertion at Beginning
Inserting the node the at the beginning of linked list
Time : O(1)
Space: O(1)
*/

package linkedlist

import (
	"errors"
	"fmt"
)

// ================================
// ListNode Definition
// ================================

type ListNode struct {
	Data any
	Next *ListNode
}

// ================================
// LinkedList Definition
// ================================

type LinkedList struct {
	Head *ListNode
	Size int
}

// ================================
// Traversal
// ================================

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

// ================================
// Insertion at Beginning
// ================================

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
