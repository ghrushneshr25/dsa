// @index: 1
// @problem: Find Kth Node from End of Linked List
// @difficulty: Easy
// @tags: linked-list, two-pointers, hashmap
// @time: O(n)
// @space: O(1) / O(n)

/*
@section: Description
Given a singly linked list, find the kth node from the end of the list.

@section: Goal
Return the data of the kth node from the end.

@section: Example
List: 1 -> 2 -> 3 -> 4 -> 5
k = 2
Output: 4

@section: Constraints
1. The list is singly linked (forward traversal only)
2. k is a positive integer; handle empty list and k larger than length as errors
3. Multiple strategies: brute force, hash map, or two pointers

@section: Approach 1 - Brute Force

@section: Algorithm
KthFromEnd_BruteForce(list, k):
  if list is empty -> error
  if k <= 0 -> error

  length <- 0
  current <- list.head
  while current != nil:
    length <- length + 1
    current <- current.next

  if k > length -> error

  target <- length - k
  current <- list.head
  repeat target times:
    current <- current.next

  return current.data

@section: Notes
- Two traversals: first count nodes, second walk (length - k) steps from the head
- Time O(n), extra space O(1)

@section: Approach 2 - Hash Map

@section: Algorithm
KthFromEnd_HashMap(list, k):
  if list is empty -> error
  if k <= 0 -> error

  map <- empty index to node
  index <- 0
  current <- list.head
  while current != nil:
    map[index] <- current
    current <- current.next
    index <- index + 1

  if k > index -> error

  return map[index - k].data

@section: Notes
- One pass to fill the map; answer is the node at index (n - k) with 0-based indices
- Time O(n), extra space O(n) for the map

@section: Approach 3 - Efficient (Two Pointer)

@section: Algorithm
KthFromEnd_TwoPointer(list, k):
  if list is empty -> error
  if k <= 0 -> error

  fast <- list.head
  slow <- list.head

  repeat k times:
    if fast is nil -> error (k exceeds list size)
    fast <- fast.next

  while fast != nil:
    fast <- fast.next
    slow <- slow.next

  return slow.data

@section: Notes
- Single pass after advancing fast by k: gap between slow and fast stays k nodes until fast exits the list
- Time O(n), extra space O(1); preferred when memory matters

@section: Summary
- Most optimal approach is Two Pointer
- Avoid extra traversal and O(n) extra space when you can use two pointers
*/

package linkedlist

import (
	"errors"
	"fmt"
)

// @operation KthFromEnd_BruteForce
// @description Uses two traversals to find kth node from end.
// @time: O(n)
// @space: O(1)
func KthFromEnd_BruteForce(list *LinkedList, k int) (any, error) {
	if list == nil || list.Head == nil {
		return nil, fmt.Errorf("list is empty")
	}

	if k <= 0 {
		return nil, fmt.Errorf("invalid k")
	}

	current := list.Head
	size := 0

	for current != nil {
		current = current.Next
		size++
	}

	if k > size {
		return nil, fmt.Errorf("k exceeds list size")
	}

	position := size - k

	current = list.Head

	for i := 0; i < position; i++ {
		current = current.Next
	}

	return current.Data, nil
}

// @operation KthFromEnd_HashMap
// @description Stores nodes in map for direct access.
// @time: O(n)
// @space: O(n)
func KthFromEnd_HashMap(list *LinkedList, k int) (any, error) {
	if list == nil || list.Head == nil {
		return nil, fmt.Errorf("list is empty")
	}

	if k <= 0 {
		return nil, fmt.Errorf("invalid k")
	}

	nodeMap := make(map[int]*ListNode)

	current := list.Head
	index := 0
	for current != nil {
		nodeMap[index] = current
		current = current.Next
		index++
	}

	if k > index {
		return nil, fmt.Errorf("k exceeds list size")
	}

	return nodeMap[index-k].Data, nil

}

// @operation KthFromEnd_TwoPointer
// @description Uses two pointers to find kth node in one traversal.
// @time: O(n)
// @space: O(1)
func KthFromEnd_TwoPointer(list *LinkedList, k int) (data any, err error) {
	if list == nil || list.Head == nil {
		return nil, errors.New("list is empty")
	}

	if k <= 0 {
		return nil, fmt.Errorf("invalid k")
	}

	slowPointer, fastPointer := list.Head, list.Head

	for i := 0; i < k; i++ {
		if fastPointer == nil {
			return nil, fmt.Errorf("k exceeds list size")
		}
		fastPointer = fastPointer.Next
	}

	for fastPointer != nil {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next
	}

	return slowPointer.Data, nil
}
