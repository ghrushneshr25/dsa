// @problem: Generate K-ary Strings
// @difficulty: Easy
// @tags: recursion, backtracking
// @time: O(k^n)
// @space: O(n)

/*
@section: Description
Generate all possible strings of length n using characters from 0 to k-1.

@section: Goal
Generate all k^n combinations where each position can take values from 0 to k-1.

@section: Constraints
1. Each position can take values in range [0, k-1]
2. Use recursion/backtracking
3. Generate all combinations

@section: Algorithm
GenerateKaryStrings(arr, index, n, k):

if index == n

	process current array
	return

for i from 0 to k-1

	arr[index] = i
	GenerateKaryStrings(arr, index+1, n, k)

@section: Notes
- Total combinations = k^n
- Generalization of binary string generation
- Forms a k-ary recursion tree
*/
package recursionbacktracking

import (
	"fmt"
	"strings"
)

func GenerateKaryStrings(array []string, index int, n int, k int, result *[]string) {
	if index == n {
		*result = append(*result, strings.Join(array, ""))
		return
	}

	for i := range k {
		array[index] = fmt.Sprintf("%d", i)
		GenerateKaryStrings(array, index+1, n, k, result)
	}
}
