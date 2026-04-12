// @problem: Check if Array is Sorted Using Recursion
// @difficulty: Easy
// @tags: recursion
// @time: O(n)
// @space: O(n)

/*
@section: Description
Given an array of elements, determine whether the array is sorted in non-decreasing order using a recursive approach.

@section: Initial Setup
- An input array of size 'n' is provided.
- The array may contain integers (or comparable elements).

@section: Objective
Check if the given array is sorted in non-decreasing order (i.e., each element is less than or equal to the next element).

@section: Constraints
1. Only recursion should be used (no iterative loops).
2. Compare adjacent elements to determine order.
3. The function should return a boolean result.

@section: Goal
Return true if the array is sorted, otherwise return false.

@section: Algorithm
Function: IsSorted(arr, n)

Base Case:
If n == 0 or n == 1 → return true

Recursive Case:
If arr[n-1] < arr[n-2] → return false
Else → IsSorted(arr, n-1)

@section: Notes
- Demonstrates recursion replacing iteration
- Can also be solved using forward traversal
*/
package recursionbacktracking

func IsArraySortedRecursion(array []int, n int) bool {
	if n == 0 || n == 1 {
		return true
	}

	if array[n-1] < array[n-2] {
		return false
	}

	return IsArraySortedRecursion(array, n-1)
}
