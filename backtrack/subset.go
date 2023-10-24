package backtrack

import (
	"fmt"
	"sort"
)

func subsets(nums []int) [][]int {
	// 子集，包括空集
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}
	// dfs 用start表示当前的起始位置，当前开始只能往后选择
	// 和组合不一样的唯一区别，就是这个会把所有的中间路径都加入到结果中，而组合需要选择全部
	var dfs func(start int, path []int)
	dfs = func(start int, path []int) {
		fmt.Println(start, path)
		// 先把上一次的结果加进来
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		for i := start; i < len(nums); i++ {
			// 遍历当前节点
			path = append(path, nums[i])
			dfs(i+1, path)
			path = path[:len(path)-1]
		}
	}
	dfs(0, []int{})
	return res
}

func subsetsWithDup(nums []int) [][]int {
	// 可以重复，则先排序。
	// 用路径的思想，如果前一个一样的没有被访问，则不会加入到当前的path中
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}
	tmp := make(Tmp, len(nums))
	copy(tmp, nums)
	sort.Sort(tmp)
	indexMap := make(map[int]bool)
	for i := range tmp {
		indexMap[i] = false
	}
	// 选择了前面节点，且当前节点与之前节点相同，则跳过，因为前面的子集已经包括了所有可能
	// 换个思路，其实是选path，path的开头不确定，所以要for each
	var dfs func(start int, path []int)
	dfs = func(start int, path []int) {
		fmt.Println(start, path, indexMap)
		// 直接加入上一轮的path
		t := make([]int, len(path))
		copy(t, path)
		res = append(res, t)

		for i := start; i < len(tmp); i++ {
			if !indexMap[i] {
				if i > 0 && !indexMap[i-1] && tmp[i] == tmp[i-1] {
					continue
				}
				indexMap[i] = true
				path = append(path, tmp[i])
				dfs(i+1, path)
				indexMap[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0, []int{})
	return res
}
