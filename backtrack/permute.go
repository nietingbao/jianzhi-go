package backtrack

import "sort"

func permute(nums []int) [][]int {
	//无重复数字，全排列
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}
	// dfs
	visitedMap := make(map[int]bool)
	for _, num := range nums {
		visitedMap[num] = false
	}
	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if !visitedMap[nums[i]] {
				visitedMap[nums[i]] = true
				path = append(path, nums[i])
				dfs(path)
				path = path[:len(path)-1]
				visitedMap[nums[i]] = false
			}
		}
	}
	dfs([]int{})
	return res
}

func permuteUnique(nums []int) [][]int {
	// 有重复数字，但是结果不能重复
	// 不重复首先想到排序，先对目标进行排序
	// dfs传入一个start参数，表示该轮是从哪个地方开始的，
	// i表示当前start轮已经进行到的dfs的位置
	// 1 1 2不能重复的话，从第1轮开始，选择第一个1，和第二个1，i是不一样的
	// 区别在于第二个i> start。
	tmp := make(Tmp, len(nums))
	copy(tmp, nums)
	res := [][]int{}
	if len(tmp) == 0 {
		return res
	}

	sort.Sort(tmp)
	//dfs 画图
	// 由于重复，vis需要加下标
	// dfs 的depth表示当前的深度，i表示现在所处的位置
	// 如果i>depth表示在选择的时候跳过了之前的数字，而如果这时该数字与前一个数字相等，说明有重复，可以跳过。
	indexMap := make(map[int]bool)
	var dfs func(depth int, path []int)
	dfs = func(depth int, path []int) {
		if depth == len(tmp) {
			t := make([]int, len(path))
			copy(t, path)
			res = append(res, t)
		}
		for i := 0; i < len(tmp); i++ {
			if !indexMap[i] {
				// 选择i
				if i > 0 && tmp[i] == tmp[i-1] && !indexMap[i-1] {
					continue
				}
				indexMap[i] = true
				path = append(path, tmp[i])
				dfs(depth+1, path)
				indexMap[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0, []int{})
	return res
}
