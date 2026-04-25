// @test: Generate Binary Strings Tests
// @covers: recursion, backtracking, correctness, edge-cases, performance

/*
@section: Overview
This test suite validates whether the recursive implementation correctly
generates all binary strings of length n using an accumulator-based approach.

It covers:
- Correct number of combinations (2^n)
- Exact sequence validation for small inputs
- Edge cases (n = 0)
- Structural correctness of generated strings
- Uniqueness of results
*/

package recursionbacktracking_test

import (
	"testing"

	recursionbacktracking "godsa/recursion-backtracking"
)

func TestGenerateBinaryStrings(t *testing.T) {

	// helper to invoke your function cleanly
	run := func(n int) []string {
		if n == 0 {
			return nil
		}
		arr := make([]string, n)
		var result []string
		recursionbacktracking.GenerateBinaryStrings(n, arr, 0, &result)
		return result
	}

	t.Run("Count Validation", func(t *testing.T) {
		// @desc: Ensures total combinations = 2^n

		tests := []struct {
			name string
			n    int
		}{
			{"n=1", 1},
			{"n=2", 2},
			{"n=3", 3},
			{"n=4", 4},
		}

		for _, tt := range tests {
			tt := tt

			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				result := run(tt.n)
				expected := 1 << tt.n

				if len(result) != expected {
					t.Fatalf("expected %d combinations, got %d", expected, len(result))
				}
			})
		}
	})

	t.Run("Exact Sequence for n=2", func(t *testing.T) {
		// @desc: Validates deterministic ordering
		t.Parallel()

		result := run(2)
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

	t.Run("Edge Case (empty length)", func(t *testing.T) {
		// @desc: Validates behavior for n = 0
		t.Parallel()

		result := run(0)

		if result != nil && len(result) != 0 {
			t.Fatalf("expected nil or empty result, got %v", result)
		}
	})

	t.Run("Binary Validity", func(t *testing.T) {
		// @desc: Ensures only '0' and '1' are present
		t.Parallel()

		result := run(3)

		for _, str := range result {
			for _, ch := range str {
				if ch != '0' && ch != '1' {
					t.Fatalf("invalid character found in %s", str)
				}
			}
		}
	})

	t.Run("Length Validation", func(t *testing.T) {
		// @desc: Ensures each string has length n
		t.Parallel()

		n := 3
		result := run(n)

		for _, str := range result {
			if len(str) != n {
				t.Fatalf("expected length %d, got %d in %s", n, len(str), str)
			}
		}
	})

	t.Run("Uniqueness Validation", func(t *testing.T) {
		// @desc: Ensures no duplicate strings exist
		t.Parallel()

		result := run(4)

		seen := make(map[string]bool)

		for _, str := range result {
			if seen[str] {
				t.Fatalf("duplicate string found: %s", str)
			}
			seen[str] = true
		}
	})
}
