package sort_algo

func duplicate(numbers []int) int {
	// write code here
	// 数字是0- n-1之内，所以用一个数组保存每一个数字是否为1即可。如果用map，可能会溢出
	l := len(numbers)
	if l == 0 {
		return -1
	}

	m := make([]int, l, l)
	for i := range m {
		m[i] = -1
	}
	for _, num := range numbers {
		old := m[num]
		if old == -1 {
			m[num] = 1
		} else {
			return num
		}
	}
	return -1
}
