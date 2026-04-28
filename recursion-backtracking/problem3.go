// @index: 3
// @problem: Generate Binary Strings
// @difficulty: Easy
// @tags: recursion, backtracking
// @time: O(2^n)
// @space: O(n)

/*
@section: Description
Generate all possible binary strings of length n.
Each position can take either 0 or 1.

@section: Goal
Generate all 2^n combinations of binary strings.

@section: Constraints
1. Each position can be 0 or 1
2. Use recursion/backtracking
3. Generate all combinations

@section: Algorithm
GenerateBinaryStrings(arr, index, n):

if index == n

	process current array
	return

arr[index] = 0
GenerateBinaryStrings(arr, index+1, n)

arr[index] = 1
GenerateBinaryStrings(arr, index+1, n)

@section: Notes
- Total combinations = 2^n
- Forms a binary recursion tree
*/

package recursionbacktracking

import "strings"

func GenerateBinaryStrings(n int, array []string, i int, result *[]string) {
	if i == n {
		*result = append(*result, strings.Join(array, ""))
		return
	}
	array[i] = "0"
	GenerateBinaryStrings(n, array, i+1, result)
	array[i] = "1"
	GenerateBinaryStrings(n, array, i+1, result)
}
