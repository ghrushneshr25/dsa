// @problem: Towers of Hanoi
// @difficulty: Easy
// @tags: recursion, backtracking, divide-and-conquer
// @time: O(2^n)
// @space: O(n)

/*
@section: Description
The Towers of Hanoi is a classic recursive problem involving three rods and multiple disks of different sizes.

All disks start on the source rod in ascending order (smallest at the top). The goal is to move all disks to the destination rod using an auxiliary rod, while following strict movement rules.

@section: Goal
Move all disks from the source rod to the destination rod such that:
- Only one disk is moved at a time
- Only the topmost disk can be moved
- No disk is placed on top of a smaller disk

@section: Constraints
1. Only one disk can be moved at a time
2. A disk must always be taken from the top
3. A larger disk cannot be placed on a smaller disk

@section: Algorithm
TowerOfHanoi(n, source, destination, auxiliary):

	if n == 1
		move disk 1 from source → destination
	return

	TowerOfHanoi(n-1, source, auxiliary, destination)

	move disk n from source → destination

	TowerOfHanoi(n-1, auxiliary, destination, source)

@section: Notes
- This problem follows a divide-and-conquer approach
- Each step breaks into smaller subproblems
- Minimum moves required = (2^n - 1)
- Widely used to understand recursion trees and call stacks
*/
package recursionbacktracking

import "fmt"

type TowerOfHanoiMove struct {
	Disk int
	From string
	To   string
}

func (m TowerOfHanoiMove) String() string {
	return fmt.Sprintf("Move disk %d from %s to %s", m.Disk, m.From, m.To)
}

func TowerOfHanoi(n int, source, destination, auxiliary string) []TowerOfHanoiMove {
	if n <= 0 {
		return nil
	}
	var moves []TowerOfHanoiMove
	if n == 1 {
		moves = append(moves, TowerOfHanoiMove{
			1, source, destination,
		})
		return moves
	}
	moves = append(moves, TowerOfHanoi(n-1, source, auxiliary, destination)...)
	moves = append(moves, TowerOfHanoiMove{n, source, destination})
	moves = append(moves, TowerOfHanoi(n-1, auxiliary, destination, source)...)
	return moves
}
