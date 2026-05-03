// @index: 5
// @problem: Largest Connected Region in Matrix
// @difficulty: Medium
// @tags: recursion, dfs, graph
// @time: O(n * m)
// @space: O(n * m)

/*
@section: Description
Given a 2D matrix of 0s and 1s, find the size of the largest connected region of 1s.
Cells are connected in 8 directions (horizontal, vertical, diagonal).

@section: Goal
Return the maximum number of cells in any connected region of 1s.

@section: Constraints
1. Each cell can be visited only once
2. Connections include 8 directions
3. Matrix is modified in-place to mark visited cells

@section: Algorithm
FindMaxConnects(matrix, M, N):

maxRegion = 0

for r from 0 to M-1
	for c from 0 to N-1

		if matrix[r][c] == 1
			size = FindConnects(matrix, M, N, r, c)
			maxRegion = max(maxRegion, size)

return maxRegion


FindConnects(matrix, M, N, r, c):

if r < 0 OR c < 0 OR r >= M OR c >= N
	return 0

if matrix[r][c] == 0
	return 0

// mark visited
matrix[r][c] = 0

count = 1

count += FindConnects(matrix, M, N, r-1, c)     // up
count += FindConnects(matrix, M, N, r+1, c)     // down
count += FindConnects(matrix, M, N, r, c-1)     // left
count += FindConnects(matrix, M, N, r, c+1)     // right

count += FindConnects(matrix, M, N, r-1, c-1)   // top-left
count += FindConnects(matrix, M, N, r-1, c+1)   // top-right
count += FindConnects(matrix, M, N, r+1, c-1)   // bottom-left
count += FindConnects(matrix, M, N, r+1, c+1)   // bottom-right

return count

@section: Notes
- Uses DFS to explore each connected component
- Matrix is modified in-place to avoid extra space
- Each cell is visited at most once
- Similar to flood fill / number of islands problem
*/

package recursionbacktracking

func LargestConnectedRegionInMatrix(matrix [][]int) int {
	rows := len(matrix)
	if rows == 0 {
		return 0
	}
	cols := len(matrix[0])

	result := 0

	var dfs func(matrix [][]int, rows, cols, currentRow, currentCol int) int

	dfs = func(matrix [][]int, rows, cols, currentRow, currentCol int) int {
		if currentRow < 0 || currentCol < 0 || currentRow >= rows || currentCol >= cols {
			return 0
		}

		if matrix[currentRow][currentCol] == 0 {
			return 0
		}

		matrix[currentRow][currentCol] = 0

		count := 1

		count += dfs(matrix, rows, cols, currentRow-1, currentCol)
		count += dfs(matrix, rows, cols, currentRow+1, currentCol)
		count += dfs(matrix, rows, cols, currentRow, currentCol-1)
		count += dfs(matrix, rows, cols, currentRow, currentCol+1)
		count += dfs(matrix, rows, cols, currentRow-1, currentCol-1)
		count += dfs(matrix, rows, cols, currentRow-1, currentCol+1)
		count += dfs(matrix, rows, cols, currentRow+1, currentCol-1)
		count += dfs(matrix, rows, cols, currentRow+1, currentCol+1)

		return count
	}

	for currentRow := 0; currentRow < rows; currentRow++ {
		for currentCol := 0; currentCol < cols; currentCol++ {
			if matrix[currentRow][currentCol] == 1 {
				size := dfs(matrix, rows, cols, currentRow, currentCol)
				if size > result {
					result = size
				}
			}
		}
	}

	return result
}
