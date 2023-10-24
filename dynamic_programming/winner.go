package dynamicprogramming

func predictTheWinner(nums []int) bool {
	// 两个玩家，输入一个数组，每一次只能从两端选择一个数，问先手是否能赢
	// d[i][j] 表示从i到j，当前玩家领先对面的分数
	// d[i][i] = nums[i]
	return false
}

func stoneGame(piles []int) bool {
	// d[i][j] 表示从i到j，
	d := make(map[int]map[int]int)
	l := len(piles)
	for i := 0; i < l; i++ {
		d[i] = map[int]int{}
		d[i][i] = piles[i]
	}
	// d[i][j] = max(piles[i] - d[i+1][j],  piles[j] - d[i][j-1])
	// i > j
	for i := l - 2; i >= 0; i-- {
		for j := i + 1; j < l; j++ {
			pickI := piles[i] - d[i+1][j]
			pickJ := piles[j] - d[i][j-1]
			if pickI > pickJ {
				d[i][j] = pickI
			} else {
				d[i][j] = pickJ
			}
		}
	}
	return d[0][l-1] > 0
}
