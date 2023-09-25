package backtrack

func movingCount(threshold int, rows int, cols int) int {
	// write code here
	if threshold == 0 {
		return 0
	}
	v := map[int]map[int]bool{}
	for i := 0; i < rows; i++ {
		vv, ok := v[i]
		if !ok {
			vv = map[int]bool{}
		}
		for j := 0; j < cols; j++ {
			vv[j] = false
		}
		v[i] = vv
	}
	// dfs 表示是否能走这一步，如果能走，则结果加一，而且会继续往东南西北走
	res := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		// 边界
		if i < 0 || i >= rows || j < 0 || j >= cols {
			return
		}
		visited := v[i][j]
		if visited {
			return
		}
		if getTotal(i)+getTotal(j) > threshold {
			return
		}
		v[i][j] = true
		res++
		dfs(i, j+1)
		dfs(i, j-1)
		dfs(i+1, j)
		dfs(i-1, j)
	}
	dfs(0, 0)
	return res
}

func getTotal(i int) int {
	res := 0
	for i >= 10 {
		res += i % 10
		i = i / 10
	}
	res += i
	return res
}
