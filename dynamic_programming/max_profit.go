package dynamicprogramming

func maxProfit(prices []int) int {
	// write code here
	// di为在第i天交易的最大值
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	min := prices[0]
	max := 0
	for i := 1; i < len(prices); i++ {
		pro := prices[i] - min
		if pro > max {
			max = pro
		}
		if pro < 0 {
			min = prices[i]
		}
	}
	return max
}
