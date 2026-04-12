// @test: Tower of Hanoi Tests
// @covers: recursion, correctness, edge-cases

/*
@section: Overview
Validates Tower of Hanoi solution across correctness, edge cases, and simulation.
*/

package recursionbacktracking_test

import (
	"reflect"
	"testing"

	recursionbacktracking "godsa/recursion-backtracking"
)

func TestTowerOfHanoi(t *testing.T) {

	t.Run("Move Count", func(t *testing.T) {
		// @desc: Validates number of moves = (2^n - 1)

		tests := []struct {
			name string
			n    int
		}{
			{"n=0", 0},
			{"n=1", 1},
			{"n=2", 2},
			{"n=3", 3},
			{"n=4", 4},
		}

		for _, tt := range tests {
			moves := recursionbacktracking.TowerOfHanoi(tt.n, "A", "C", "B")

			expected := 0
			if tt.n > 0 {
				expected = (1 << tt.n) - 1
			}

			if len(moves) != expected {
				t.Fatalf("expected %d moves, got %d", expected, len(moves))
			}
		}
	})

	t.Run("Exact Sequence N3", func(t *testing.T) {
		// @desc: Validates exact move sequence for n=3

		moves := recursionbacktracking.TowerOfHanoi(3, "A", "C", "B")

		expected := []recursionbacktracking.TowerOfHanoiMove{
			{1, "A", "C"},
			{2, "A", "B"},
			{1, "C", "B"},
			{3, "A", "C"},
			{1, "B", "A"},
			{2, "B", "C"},
			{1, "A", "C"},
		}

		if !reflect.DeepEqual(moves, expected) {
			t.Fatalf("expected %+v, got %+v", expected, moves)
		}
	})

	t.Run("Zero Disks", func(t *testing.T) {
		// @desc: Edge case where n = 0 should return no moves

		moves := recursionbacktracking.TowerOfHanoi(0, "A", "C", "B")

		if len(moves) != 0 {
			t.Fatalf("expected no moves, got %v", moves)
		}
	})

	t.Run("Valid Moves Simulation", func(t *testing.T) {
		// @desc: Simulates stack behavior to ensure valid moves

		n := 3
		moves := recursionbacktracking.TowerOfHanoi(n, "A", "C", "B")

		stacks := map[string][]int{
			"A": {3, 2, 1},
			"B": {},
			"C": {},
		}

		for _, move := range moves {
			from := stacks[move.From]
			to := stacks[move.To]

			disk := from[len(from)-1]

			if len(to) > 0 && to[len(to)-1] < disk {
				t.Fatalf("invalid move")
			}

			stacks[move.From] = from[:len(from)-1]
			stacks[move.To] = append(to, disk)
		}
	})
}
