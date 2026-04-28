// @type: concept
// @index: 4
// @title: XOR Linked List
// @tags: linked-list, advanced, memory-optimization
// @difficulty: Hard

/*
@section: Description
An XOR Linked List is a memory-efficient version of a doubly linked list.
Instead of storing separate prev and next pointers, each node stores a single field:
    both = XOR(prev, next)

@section: Structure
Each node contains:
    Data
    Both (XOR of prev and next node addresses)

@section: Operations
Traversal, insertion, and deletion require reconstructing next/prev using XOR.
*/

package linkedlist

import (
	"errors"
	"fmt"
	"unsafe"
)

// @structure XOR List Node
// @description Represents a node in XOR linked list storing XOR of prev and next.
type XORListNode struct {
	Data any
	Both uintptr // XOR of prev and next node addresses
}

// @structure XORLinkedList
// @description Represents the XOR linked list with head pointer and size.
type XORLinkedList struct {
	Head *XORListNode
	Size int
}

// @operation XOR Helper
// @description Returns XOR of two node pointers
func (list *XORLinkedList) xor(firstNode, secondNode *XORListNode) uintptr {
	return uintptr(unsafe.Pointer(firstNode)) ^ uintptr(unsafe.Pointer(secondNode))
}

// @operation Dereference Pointer
// @description Converts uintptr back to *XORListNode
func (list *XORLinkedList) deref(addr uintptr) *XORListNode {
	return (*XORListNode)(unsafe.Pointer(addr))
}

// @operation Increase Size
func (list *XORLinkedList) IncreaseSize() {
	list.Size++
}

// @operation Decrease Size
func (list *XORLinkedList) DecreaseSize() {
	list.Size--
}

// @operation Get Size
func (list *XORLinkedList) GetSize() int {
	return list.Size
}

// @operation Traversal
// @description Traverse XOR linked list
// @time: O(n)
// @space: O(1)
func (list *XORLinkedList) Display() (result string, err error) {
	if list.Head == nil {
		err = errors.New("list is empty")
		return
	}

	var prev *XORListNode = nil
	current := list.Head

	for current != nil {
		result = fmt.Sprintf("%+v -> %+v", result, current.Data)
		nextAddress := current.Both ^ uintptr(unsafe.Pointer(prev))
		next := list.deref(nextAddress)

		prev = current
		current = next
	}
	return
}

// @operation Insert at Front
// @description Insert node at beginning
// @time: O(1)
// @space: O(1)
func (list *XORLinkedList) InsertAtFront(data any) (err error) {
	newNode := &XORListNode{Data: data}

	// Case 1 : Empty List
	if list.Head == nil {
		newNode.Both = 0
		list.Head = newNode
		list.IncreaseSize()
		return
	}

	// Case 2: Non Empty List
	oldHead := list.Head
	newNode.Both = uintptr(unsafe.Pointer(oldHead))
	oldNext := list.deref(oldHead.Both ^ uintptr(unsafe.Pointer(nil)))
	oldHead.Both = list.xor(newNode, oldNext)
	list.Head = newNode
	list.IncreaseSize()
	return nil
}

// @operation Insert at End
// @description Insert node at end
// @time: O(n)
// @space: O(1)
func (list *XORLinkedList) InsertAtEnd(data any) (err error) {
	newNode := &XORListNode{Data: data}

	// Case 1 : Empty List
	if list.Head == nil {
		newNode.Both = 0
		list.Head = newNode
		list.IncreaseSize()
		return
	}

	// Case 2: Non Empty List

	var prev *XORListNode = nil
	current := list.Head

	for {
		nextAddr := current.Both ^ uintptr(unsafe.Pointer(prev))
		next := list.deref(nextAddr)
		if next == nil {
			break
		}
		prev = current
		current = next
	}

	newNode.Both = uintptr(unsafe.Pointer(current))

	current.Both = list.xor(prev, newNode)

	list.IncreaseSize()
	return nil
}

// @operation Insert at Position
// @description Insert node at given position
// @time: O(n)
// @space: O(1)
func (list *XORLinkedList) Insert(position int, data any) (err error) {
	if position < 1 || position > list.GetSize()+1 {
		err = fmt.Errorf("invalid position")
		return
	}

	newNode := &XORListNode{Data: data}

	if list.Head == nil {
		newNode.Both = 0
		list.Head = newNode
		list.IncreaseSize()
		return
	}

	size := list.GetSize()
	position = (position-1)%size + 1

	if position == 1 {
		return list.InsertAtFront(data)
	}

	var prev *XORListNode = nil
	current := list.Head

	for i := 1; i < position; i++ {
		nextAddr := current.Both ^ uintptr(unsafe.Pointer(prev))
		next := list.deref(nextAddr)

		prev = current
		current = next
	}

	newNode.Both = list.xor(prev, current)

	prevPrevAddr := prev.Both ^ uintptr(unsafe.Pointer(current))
	prevPrev := list.deref(prevPrevAddr)

	prev.Both = list.xor(prevPrev, newNode)

	nextAddr := current.Both ^ uintptr(unsafe.Pointer(prev))
	next := list.deref(nextAddr)
	current.Both = list.xor(newNode, next)

	list.IncreaseSize()
	return
}

// @operation Delete From Front
// @description Delete node from beginning
// @time: O(1)
// @space: O(1)
func (list *XORLinkedList) DeleteFromFront() (data any, err error) {
	if list.Head == nil {
		return nil, fmt.Errorf("empty list")
	}

	current := list.Head
	data = current.Data

	// Case 1: Single node
	if list.GetSize() == 1 {
		list.Head = nil
		list.DecreaseSize()
		return data, nil
	}

	next := list.deref(current.Both)
	nextNextAddr := next.Both ^ uintptr(unsafe.Pointer(current))
	nextNext := list.deref(nextNextAddr)

	next.Both = list.xor(nil, nextNext)

	list.Head = next
	list.DecreaseSize()
	return
}

// @operation Delete From End
// @description Delete node from end
// @time: O(n)
// @space: O(1)
func (list *XORLinkedList) DeleteFromEnd() (data any, err error) {
	if list.Head == nil {
		err = fmt.Errorf("list is empty")
		return
	}

	// Case 1: Single node
	if list.GetSize() == 1 {
		data := list.Head.Data
		list.Head = nil
		list.DecreaseSize()
		return data, nil
	}

	var prev *XORListNode = nil
	current := list.Head

	for {
		nextAddr := current.Both ^ uintptr(unsafe.Pointer(prev))
		next := list.deref(nextAddr)

		if next == nil {
			break
		}

		prev = current
		current = next
	}

	data = current.Data

	prevPrevAddr := prev.Both ^ uintptr(unsafe.Pointer(current))
	prevPrev := list.deref(prevPrevAddr)

	prev.Both = list.xor(prevPrev, nil)
	list.DecreaseSize()
	return
}

// @operation Delete From Position
// @description Delete node at given position
// @time: O(n)
// @space: O(1)
func (list *XORLinkedList) Delete(position int) (data any, err error) {
	if list.Head == nil {
		return nil, fmt.Errorf("empty list")
	}

	if position <= 0 {
		return nil, fmt.Errorf("invalid position")
	}

	size := list.GetSize()
	position = (position-1)%size + 1

	// Edge cases
	if position == 1 {
		return list.DeleteFromFront()
	}
	if position == size {
		return list.DeleteFromEnd()
	}

	var prev *XORListNode = nil
	current := list.Head

	for i := 1; i < position; i++ {
		nextAddr := current.Both ^ uintptr(unsafe.Pointer(prev))
		next := list.deref(nextAddr)

		prev = current
		current = next
	}

	data = current.Data
	// Get next
	nextAddr := current.Both ^ uintptr(unsafe.Pointer(prev))
	next := list.deref(nextAddr)

	// Get prevPrev
	prevPrevAddr := prev.Both ^ uintptr(unsafe.Pointer(current))
	prevPrev := list.deref(prevPrevAddr)

	// Get nextNext
	nextNextAddr := next.Both ^ uintptr(unsafe.Pointer(current))
	nextNext := list.deref(nextNextAddr)

	// Fix prev
	prev.Both = list.xor(prevPrev, next)

	// Fix next
	next.Both = list.xor(prev, nextNext)

	list.DecreaseSize()

	return data, nil
}
