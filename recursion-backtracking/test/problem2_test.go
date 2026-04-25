// @test: Is Array Sorted (Recursion) Tests
// @covers: recursion, correctness, edge-cases, performance

/*
@section: Overview
This test suite validates whether the recursive implementation correctly determines
if an array is sorted in non-decreasing order.

It covers:
- Basic sorted/unsorted cases
- Edge cases (empty, single element)
- Negative numbers
- Large input validation
- Early violation detection
*/

package recursionbacktracking_test

import (
	"testing"

	recursionbacktracking "godsa/recursion-backtracking"
)

func TestIsArraySortedRecursion(t *testing.T) {

	t.Run("Basic Cases", func(t *testing.T) {
		// @desc: Validates common sorted and unsorted scenarios

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

	t.Run("Edge Cases", func(t *testing.T) {
		// @desc: Validates boundary conditions like empty and single element arrays

		t.Parallel()

		if !recursionbacktracking.IsArraySortedRecursion([]int{}, 0) {
			t.Fatalf("expected true for empty array")
		}

		if !recursionbacktracking.IsArraySortedRecursion([]int{10}, 1) {
			t.Fatalf("expected true for single element")
		}
	})

	t.Run("Negative Numbers", func(t *testing.T) {
		// @desc: Ensures sorting works correctly with negative numbers

		t.Parallel()

		input := []int{-5, -3, -1, 0, 2}
		result := recursionbacktracking.IsArraySortedRecursion(input, len(input))

		if !result {
			t.Fatalf("expected true for sorted negative numbers, got false")
		}
	})

	t.Run("Large Input", func(t *testing.T) {
		// @desc: Tests performance and correctness for large input sizes

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

	t.Run("First Violation Detection", func(t *testing.T) {
		// @desc: Detects first violation in ordering

		t.Parallel()

		input := []int{1, 2, 3, 7, 5, 6}
		result := recursionbacktracking.IsArraySortedRecursion(input, len(input))

		if result {
			t.Fatalf("expected false due to violation at index 3->4")
		}
	})
}
