package backtrack

func hasPath(matrix [][]byte, word string) bool {
	// write code here
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}
	if len(matrix)*len(matrix[0]) < len(word) {
		return false
	}
	v := map[int]map[int]bool{}
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		vv, ok := v[i]
		if !ok {
			vv = map[int]bool{}
		}
		for j := 0; j < n; j++ {
			vv[j] = false
		}
		v[i] = vv
	}

	var dfs func(i, j int, index int) bool
	// dfs 的内部调用肯定会使用参数的变化形态。
	dfs = func(i, j int, index int) bool {
		if i < 0 || i >= m || j < 0 || j >= n {
			return false
		}
		visited := v[i][j]
		if visited {
			return false
		}
		if matrix[i][j] != word[index] {
			return false
		}
		v[i][j] = true
		if index == len(word)-1 {
			return true
		}
		if dfs(i, j+1, index+1) || dfs(i, j-1, index+1) || dfs(i+1, j, index+1) || dfs(i-1, j, index+1) {
			return true
		}
		v[i][j] = false
		return false
	}
	// dfs是，基于当前的结果，继续调用下面的结果
	// 描述当前的状态
	// word的字母索引
	index := 0
	// 需要在所有的位置开始找，找到一个即从内部调用dfs
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, index) {
				return true
			}
		}
	}
	return false
}
