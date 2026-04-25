// @test: Generate K-ary Strings Tests
// @covers: recursion, backtracking, correctness, edge-cases, performance

/*
@section: Overview
This test suite validates whether the recursive implementation correctly
generates all k-ary strings of length n.

It covers:
- Correct number of combinations (k^n)
- Exact sequence validation for small inputs
- Edge cases (n = 0, k = 0)
- Structural correctness of generated strings
- Uniqueness of results
*/

package recursionbacktracking_test

import (
	"testing"

	recursionbacktracking "godsa/recursion-backtracking"
)

func TestGenerateKaryStrings(t *testing.T) {

	// helper wrapper
	run := func(n, k int) []string {
		if n == 0 || k == 0 {
			return nil
		}
		arr := make([]string, n)
		var result []string
		recursionbacktracking.GenerateKaryStrings(arr, 0, n, k, &result)
		return result
	}

	t.Run("Count Validation", func(t *testing.T) {
		// @desc: Ensures total combinations = k^n

		tests := []struct {
			name string
			n    int
			k    int
		}{
			{"n=1,k=2", 1, 2},
			{"n=2,k=2", 2, 2},
			{"n=2,k=3", 2, 3},
			{"n=3,k=2", 3, 2},
		}

		for _, tt := range tests {
			tt := tt

			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				result := run(tt.n, tt.k)

				expected := 1
				for i := 0; i < tt.n; i++ {
					expected *= tt.k
				}

				if len(result) != expected {
					t.Fatalf("expected %d combinations, got %d", expected, len(result))
				}
			})
		}
	})

	t.Run("Exact Sequence n=2 k=2", func(t *testing.T) {
		// @desc: Validates deterministic ordering (binary case)
		t.Parallel()

		result := run(2, 2)
		expected := []string{"00", "01", "10", "11"}

		if len(result) != len(expected) {
			t.Fatalf("expected %v, got %v", expected, result)
		}

		for i := range expected {
			if result[i] != expected[i] {
				t.Fatalf("expected %v, got %v", expected, result)
			}
		}
	})

	t.Run("ExactSequence n=2 k=3", func(t *testing.T) {
		// @desc: Validates ordering for k=3
		t.Parallel()

		result := run(2, 3)
		expected := []string{
			"00", "01", "02",
			"10", "11", "12",
			"20", "21", "22",
		}

		if len(result) != len(expected) {
			t.Fatalf("expected %v, got %v", expected, result)
		}

		for i := range expected {
			if result[i] != expected[i] {
				t.Fatalf("expected %v, got %v", expected, result)
			}
		}
	})

	t.Run("Edge Cases", func(t *testing.T) {
		t.Parallel()

		if res := run(0, 3); len(res) != 0 {
			t.Fatalf("expected empty for n=0, got %v", res)
		}

		if res := run(3, 0); len(res) != 0 {
			t.Fatalf("expected empty for k=0, got %v", res)
		}
	})

	t.Run("Length Validation", func(t *testing.T) {
		t.Parallel()

		n, k := 3, 3
		result := run(n, k)

		for _, str := range result {
			if len(str) != n {
				t.Fatalf("expected length %d, got %d in %s", n, len(str), str)
			}
		}
	})

	t.Run("RangeValidation", func(t *testing.T) {
		// @desc: Ensures digits are within [0, k-1]
		t.Parallel()

		n, k := 3, 3
		result := run(n, k)

		for _, str := range result {
			for _, ch := range str {
				if ch < '0' || ch >= rune('0'+k) {
					t.Fatalf("invalid character %c in %s", ch, str)
				}
			}
		}
	})

	t.Run("Uniqueness Validation", func(t *testing.T) {
		t.Parallel()

		result := run(3, 3)
		seen := make(map[string]bool)

		for _, str := range result {
			if seen[str] {
				t.Fatalf("duplicate string found: %s", str)
			}
			seen[str] = true
		}
	})

	t.Run("Large Input", func(t *testing.T) {
		t.Parallel()

		n, k := 6, 2
		result := run(n, k)

		expected := 1 << n // since k=2

		if len(result) != expected {
			t.Fatalf("expected %d combinations, got %d", expected, len(result))
		}
	})
}
