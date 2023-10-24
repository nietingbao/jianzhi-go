package backtrack

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	// candidates 是一串数字，从里面找到满足target 的一个路径
	// 数字可以重新选择, 选择一个之后，可以递归继续选择
	// 选完之后还是不对，则回溯
	// 还有个问题，如果是结果一样，只是顺序不一样的路径，如何去重？
	// 跟组合一样，每次只能从当前以及后面选择， 这样就不会出现重复了
	// 加一个参数即可
	if len(candidates) == 0 {
		return nil
	}
	res := [][]int{}
	var dfs func(start int, path []int, target int)
	dfs = func(start int, path []int, target int) {
		if target == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		if target < 0 {
			return
		}
		// 可能从任意一个位置开始选择
		for i := start; i < len(candidates); i++ {
			//选择i
			path = append(path, candidates[i])
			//继续从头递归选择
			dfs(i, path, target-candidates[i])
			// 回溯
			path = path[:len(path)-1]
		}
	}
	dfs(0, []int{}, target)
	return res
}

type Tmp []int

func (t Tmp) Len() int {
	return len(t)
}

func (t Tmp) Less(i, j int) bool {
	return t[i] < t[j]
}

func (t Tmp) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func combinationSum2(candidates []int, target int) [][]int {
	// candidates 是一串数字，从里面找到满足target 的一个路径
	// 数字可以重新选择, 选择一个之后，可以递归继续选择
	// 选完之后还是不对，则回溯
	// 还有个问题，如果是结果一样，只是顺序不一样的路径，如何去重？
	// 跟组合一样，每次只能从当前以及后面选择， 这样就不会出现重复了
	// 加一个参数即可
	tmp := make(Tmp, len(candidates))
	copy(tmp, candidates)
	sort.Sort(tmp)
	if len(tmp) == 0 {
		return nil
	}
	visited := make(map[int]bool)
	for can := range tmp {
		visited[can] = false
	}
	res := [][]int{}
	var dfs func(start int, path []int, target int)
	dfs = func(start int, path []int, target int) {
		// 对最初或者上一轮的判断
		if target == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		if target < 0 {
			return
		}

		// 可能从任意一个位置开始选择
		for i := start; i < len(tmp); i++ {
			// 为啥会有1 1 6但1 2 5 不重复？
			if i > 0 && !visited[i-1] && tmp[i] == tmp[i-1] {
				continue
			}
			//选择i
			path = append(path, tmp[i])
			visited[i] = true
			//从后面选择，就可以不判断visited了
			dfs(i+1, path, target-tmp[i])
			// 回溯
			path = path[:len(path)-1]
			visited[i] = false
		}
	}
	dfs(0, []int{}, target)
	return res
}
