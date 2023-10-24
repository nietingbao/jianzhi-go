package dynamicprogramming

func coinChange(coins []int, amount int) int {
	if len(coins) == 0 || amount == 0 {
		return 0
	}
	// d[i] = min(for j in coins; d[i-coins[j]] + 1)
	d := make(map[int]int)
	for i := 0; i <= amount; i++ {
		d[i] = -1
	}
	for i := 1; i <= amount; i++ {
		min := 1<<63 - 1
		for j, coin := range coins {
			if coin == i {
				d[i] = 1
				break
			} else {
				if i-coins[j] > 0 {
					if d[i-coins[j]]+1 < min {
						if d[i-coins[j]] != -1 {
							min = d[i-coins[j]] + 1
						}

					}
				}
			}
		}
		if min != 1<<63-1 && d[i] != 1 {
			d[i] = min
		}
	}
	return d[amount]
}
