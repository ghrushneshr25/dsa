// @test: Kth Node From End Tests
// @covers: brute-force, hashmap, two-pointer, edge-cases, consistency

/*
@section: Overview
Validates all approaches for finding kth node from end across:
- correctness
- edge cases
- consistency between implementations
*/

package linkedlist_test

import (
	"testing"

	linkedlist "godsa/linkedlist"
)

func buildList(values []int) *linkedlist.LinkedList {
	list := &linkedlist.LinkedList{}
	for i := len(values) - 1; i >= 0; i-- {
		list.InsertAtBeginning(values[i])
	}
	return list
}

//
// ==========================
// 🧪 Tests
// ==========================
//

func TestKthFromEnd(t *testing.T) {

	t.Run("Basic Correctness", func(t *testing.T) {
		// @desc: Validates expected outputs for normal cases

		list := buildList([]int{1, 2, 3, 4, 5})

		tests := []struct {
			k        int
			expected int
		}{
			{1, 5},
			{2, 4},
			{3, 3},
			{5, 1},
		}

		for _, tt := range tests {
			result, _ := linkedlist.KthFromEnd_TwoPointer(list, tt.k)
			if result != tt.expected {
				t.Fatalf("k=%d expected %d got %v", tt.k, tt.expected, result)
			}

			result, _ = linkedlist.KthFromEnd_BruteForce(list, tt.k)
			if result != tt.expected {
				t.Fatalf("k=%d expected %d got %v", tt.k, tt.expected, result)
			}

			result, _ = linkedlist.KthFromEnd_HashMap(list, tt.k)
			if result != tt.expected {
				t.Fatalf("k=%d expected %d got %v", tt.k, tt.expected, result)
			}
		}
	})

	t.Run("All Approaches Consistency", func(t *testing.T) {
		// @desc: Ensures all 3 approaches return same result

		list := buildList([]int{10, 20, 30, 40, 50})

		k := 2

		bf, _ := linkedlist.KthFromEnd_BruteForce(list, k)
		hm, _ := linkedlist.KthFromEnd_HashMap(list, k)
		tp, _ := linkedlist.KthFromEnd_TwoPointer(list, k)

		if bf != hm || hm != tp {
			t.Fatalf("mismatch: brute=%v hashmap=%v twoPointer=%v", bf, hm, tp)
		}
	})

	t.Run("Single Element", func(t *testing.T) {
		// @desc: Edge case: list with one node

		list := buildList([]int{42})

		result, _ := linkedlist.KthFromEnd_TwoPointer(list, 1)
		if result != 42 {
			t.Fatalf("expected 42 got %v", result)
		}

		result, _ = linkedlist.KthFromEnd_BruteForce(list, 1)
		if result != 42 {
			t.Fatalf("expected 42 got %v", result)
		}

		result, _ = linkedlist.KthFromEnd_HashMap(list, 1)
		if result != 42 {
			t.Fatalf("expected 42 got %v", result)
		}
	})

	t.Run("Invalid K", func(t *testing.T) {
		// @desc: k less than equal to 0 should error

		list := buildList([]int{1, 2, 3})

		_, err := linkedlist.KthFromEnd_TwoPointer(list, 0)
		if err == nil {
			t.Fatalf("expected error for k=0")
		}

		_, err = linkedlist.KthFromEnd_TwoPointer(list, -1)
		if err == nil {
			t.Fatalf("expected error for negative k")
		}

		_, err = linkedlist.KthFromEnd_BruteForce(list, 0)
		if err == nil {
			t.Fatalf("expected error for k=0")
		}

		_, err = linkedlist.KthFromEnd_BruteForce(list, -1)
		if err == nil {
			t.Fatalf("expected error for negative k")
		}

		_, err = linkedlist.KthFromEnd_HashMap(list, 0)
		if err == nil {
			t.Fatalf("expected error for k=0")
		}

		_, err = linkedlist.KthFromEnd_HashMap(list, -1)
		if err == nil {
			t.Fatalf("expected error for negative k")
		}
	})

	t.Run("K Greater Than Length", func(t *testing.T) {
		// @desc: k > size should error

		list := buildList([]int{1, 2, 3})

		_, err := linkedlist.KthFromEnd_TwoPointer(list, 5)

		if err == nil {
			t.Fatalf("expected error for k > length")
		}

		_, err = linkedlist.KthFromEnd_HashMap(list, 5)

		if err == nil {
			t.Fatalf("expected error for k > length")
		}

		_, err = linkedlist.KthFromEnd_BruteForce(list, 5)

		if err == nil {
			t.Fatalf("expected error for k > length")
		}
	})

	t.Run("Empty List", func(t *testing.T) {
		// @desc: nil or empty list should error

		list := &linkedlist.LinkedList{}

		_, err := linkedlist.KthFromEnd_TwoPointer(list, 1)

		if err == nil {
			t.Fatalf("expected error for empty list")
		}

		_, err = linkedlist.KthFromEnd_HashMap(list, 5)

		if err == nil {
			t.Fatalf("expected error for empty list")
		}

		_, err = linkedlist.KthFromEnd_BruteForce(list, 5)

		if err == nil {
			t.Fatalf("expected error for empty list")
		}
	})

	t.Run("Stress Small Range", func(t *testing.T) {
		// @desc: Validates correctness across multiple inputs

		values := []int{1, 2, 3, 4, 5, 6, 7, 8}
		list := buildList(values)

		for k := 1; k <= len(values); k++ {
			expected := values[len(values)-k]

			result, err := linkedlist.KthFromEnd_TwoPointer(list, k)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result != expected {
				t.Fatalf("k=%d expected %v got %v", k, expected, result)
			}

			result, err = linkedlist.KthFromEnd_BruteForce(list, k)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result != expected {
				t.Fatalf("k=%d expected %v got %v", k, expected, result)
			}

			result, err = linkedlist.KthFromEnd_HashMap(list, k)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result != expected {
				t.Fatalf("k=%d expected %v got %v", k, expected, result)
			}
		}
	})
}
