package dynamicprogramming

func maxValue(grid [][]int) int {
	// write code here
	// d[i][j] = max(d[i-1][j] + grid[i][j], d[i][j-1] + grid[i][j])
	d := map[int]map[int]int{}
	for i := 0; i < len(grid); i++ {
		if len(d[i]) == 0 {
			d[i] = map[int]int{}
		}
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 {
				if j == 0 {
					d[i][j] = grid[0][0]
				} else {
					d[i][j] = grid[i][j] + d[0][j-1]
				}
			} else {
				if j == 0 {
					d[i][j] = grid[i][j] + d[i-1][0]
				}
			}

			if i > 0 && j > 0 {
				if d[i-1][j] > d[i][j-1] {
					d[i][j] = grid[i][j] + d[i-1][j]
				} else {
					d[i][j] = grid[i][j] + d[i][j-1]
				}
			}
		}
	}
	return d[len(grid)-1][len(grid[0])-1]
}
