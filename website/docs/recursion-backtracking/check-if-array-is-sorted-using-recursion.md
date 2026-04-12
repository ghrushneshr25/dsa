---
title: Check if Array is Sorted Using Recursion
---

**Difficulty:** Easy  
**Tags:** recursion  

**Time:** O(n)  
**Space:** O(n)  


## Description

Given an array of elements, determine whether the array is sorted in non-decreasing order using a recursive approach.



## Initial Setup

- An input array of size 'n' is provided.
- The array may contain integers (or comparable elements).



## Objective

Check if the given array is sorted in non-decreasing order (i.e., each element is less than or equal to the next element).



## Constraints

1. Only recursion should be used (no iterative loops).
2. Compare adjacent elements to determine order.
3. The function should return a boolean result.



## Goal

Return true if the array is sorted, otherwise return false.



## Algorithm

```text
Function: IsSorted(arr, n)

Base Case:
If n == 0 or n == 1 → return true

Recursive Case:
If arr[n-1] < arr[n-2] → return false
Else → IsSorted(arr, n-1)


```

## Code (Go)

```go

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
```


## Tests

### Basic Cases

Validates common sorted and unsorted scenarios

```go
	t.Run("Basic Cases", func(t *testing.T) {

		tests := []struct {
			name     string
			input    []int
			expected bool
		}{
			{"empty array", []int{}, true},
			{"single element", []int{1}, true},
			{"sorted array", []int{1, 2, 3, 4, 5}, true},
			{"unsorted array", []int{1, 3, 2, 4, 5}, false},
			{"descending array", []int{5, 4, 3, 2, 1}, false},
			{"all equal elements", []int{2, 2, 2, 2}, true},
		}

		for _, tt := range tests {
			tt := tt

			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				result := recursionbacktracking.IsArraySortedRecursion(tt.input, len(tt.input))

				if result != tt.expected {
					t.Fatalf("expected %v, got %v for input %v", tt.expected, result, tt.input)
				}
			})
		}
	})

```

### Edge Cases

Validates boundary conditions like empty and single element arrays

```go
	t.Run("Edge Cases", func(t *testing.T) {

		t.Parallel()

		if !recursionbacktracking.IsArraySortedRecursion([]int{}, 0) {
			t.Fatalf("expected true for empty array")
		}

		if !recursionbacktracking.IsArraySortedRecursion([]int{10}, 1) {
			t.Fatalf("expected true for single element")
		}
	})

```

### Negative Numbers

Ensures sorting works correctly with negative numbers

```go
	t.Run("Negative Numbers", func(t *testing.T) {

		t.Parallel()

		input := []int{-5, -3, -1, 0, 2}
		result := recursionbacktracking.IsArraySortedRecursion(input, len(input))

		if !result {
			t.Fatalf("expected true for sorted negative numbers, got false")
		}
	})

```

### Large Input

Tests performance and correctness for large input sizes

```go
	t.Run("Large Input", func(t *testing.T) {

		t.Parallel()

		n := 1000
		input := make([]int, n)
		for i := 0; i < n; i++ {
			input[i] = i
		}

		result := recursionbacktracking.IsArraySortedRecursion(input, len(input))

		if !result {
			t.Fatalf("expected true for large sorted input")
		}
	})

```

### First Violation Detection

Detects first violation in ordering

```go
	t.Run("First Violation Detection", func(t *testing.T) {

		t.Parallel()

		input := []int{1, 2, 3, 7, 5, 6}
		result := recursionbacktracking.IsArraySortedRecursion(input, len(input))

		if result {
			t.Fatalf("expected false due to violation at index 3->4")
		}
	})

```

[View Source](https://github.com/ghrushneshr25/dsa/blob/master/codebase/recursion-backtracking/problem2.go)

