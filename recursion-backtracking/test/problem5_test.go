package recursionbacktracking_test

import (
	"testing"

	recursionbacktracking "godsa/recursion-backtracking"
)

// @test: Largest Connected Region in Matrix Tests
// @covers: dfs, recursion, correctness, edge-cases, mutation-handling

/*
@section: Overview
This test suite validates whether the DFS-based implementation correctly computes
the size of the largest connected region of 1s in a matrix.

It covers:
- Basic region detection
- Multiple disconnected regions
- Diagonal connectivity
- Edge cases (empty, all 0s, all 1s)
- Single cell matrices
- Mutation safety (input modification)
*/

func TestLargestConnectedRegionInMatrix(t *testing.T) {
	copyMatrix := func(matrix [][]int) [][]int {
		rows := len(matrix)
		result := make([][]int, rows)
		for i := range matrix {
			result[i] = make([]int, len(matrix[i]))
			copy(result[i], matrix[i])
		}
		return result
	}

	t.Run("Basic Cases", func(t *testing.T) {
		t.Parallel()

		matrix := [][]int{
			{1, 1, 0},
			{0, 1, 0},
			{0, 0, 1},
		}

		result := recursionbacktracking.LargestConnectedRegionInMatrix(copyMatrix(matrix))

		if result != 4 {
			t.Fatalf("expected 3, got %d", result)
		}
	})

	t.Run("All Zeros", func(t *testing.T) {
		t.Parallel()

		matrix := [][]int{
			{0, 0},
			{0, 0},
		}

		result := recursionbacktracking.LargestConnectedRegionInMatrix(copyMatrix(matrix))

		if result != 0 {
			t.Fatalf("expected 0, got %d", result)
		}
	})

	t.Run("All Ones", func(t *testing.T) {
		t.Parallel()

		matrix := [][]int{
			{1, 1},
			{1, 1},
		}

		result := recursionbacktracking.LargestConnectedRegionInMatrix(copyMatrix(matrix))

		if result != 4 {
			t.Fatalf("expected 4, got %d", result)
		}
	})

	t.Run("Single Cell", func(t *testing.T) {
		t.Parallel()

		matrix := [][]int{{1}}

		result := recursionbacktracking.LargestConnectedRegionInMatrix(copyMatrix(matrix))

		if result != 1 {
			t.Fatalf("expected 1, got %d", result)
		}
	})

	t.Run("Single Cell Zero", func(t *testing.T) {
		t.Parallel()

		matrix := [][]int{{0}}

		result := recursionbacktracking.LargestConnectedRegionInMatrix(copyMatrix(matrix))

		if result != 0 {
			t.Fatalf("expected 0, got %d", result)
		}
	})

	t.Run("Diagonal Connection", func(t *testing.T) {
		t.Parallel()

		matrix := [][]int{
			{1, 0},
			{0, 1},
		}

		// diagonal counts as connected → size = 2
		result := recursionbacktracking.LargestConnectedRegionInMatrix(copyMatrix(matrix))

		if result != 2 {
			t.Fatalf("expected 2, got %d", result)
		}
	})

	t.Run("Multiple Regions", func(t *testing.T) {
		t.Parallel()

		matrix := [][]int{
			{1, 0, 0, 1},
			{1, 1, 0, 0},
			{0, 0, 1, 1},
		}

		result := recursionbacktracking.LargestConnectedRegionInMatrix(copyMatrix(matrix))

		if result != 5 {
			t.Fatalf("expected 3, got %d", result)
		}
	})

	t.Run("Complex Shape", func(t *testing.T) {
		t.Parallel()

		matrix := [][]int{
			{1, 1, 0, 0},
			{1, 1, 1, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 1},
		}

		result := recursionbacktracking.LargestConnectedRegionInMatrix(copyMatrix(matrix))

		if result != 8 {
			t.Fatalf("expected 6, got %d", result)
		}
	})

	t.Run("Empty Matrix", func(t *testing.T) {
		t.Parallel()

		var matrix [][]int

		result := recursionbacktracking.LargestConnectedRegionInMatrix(matrix)

		if result != 0 {
			t.Fatalf("expected 0, got %d", result)
		}
	})

	t.Run("Mutation Check", func(t *testing.T) {

		matrix := [][]int{
			{1, 1},
			{1, 0},
		}

		_ = recursionbacktracking.LargestConnectedRegionInMatrix(matrix)

		if matrix[0][0] != 0 {
			t.Fatalf("expected matrix to be mutated (visited marked)")
		}
	})
}
