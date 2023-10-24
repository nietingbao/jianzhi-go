package backtrack

import (
	"strconv"
)

func readBinaryWatch(turnedOn int) []string {
	hours := []int{1, 2, 4, 8}
	mins := []int{1, 2, 4, 8, 16, 32}
	res := []string{}
	time2str := func(hour, min int) string {
		res := strconv.Itoa(hour) + ":"
		if min < 10 {
			return res + "0" + strconv.Itoa(min)
		}
		return res + strconv.Itoa(min)
	}

	for i := 0; i <= turnedOn; i++ {
		hourParts := getNums(i, hours, 12)
		minParts := getNums(turnedOn-i, mins, 60)
		for _, hour := range hourParts {
			for _, min := range minParts {
				res = append(res, time2str(hour, min))
			}
		}
	}

	return res
}

// 组合
func getNums(turnOn int, nums []int, limit int) []int {
	if turnOn == 0 {
		return []int{0}
	}
	res := []int{}
	// 由于都是往后选的，所以不会存在选一样的
	//dfs 表示从nums中选择target个数字，获取所有的可能
	var dfs func(i int, chosed []int, tmp int)
	dfs = func(start int, chosed []int, tmp int) {
		if len(chosed) == turnOn {
			if tmp < limit {
				res = append(res, tmp)
			}
			return
		}
		for i := start; i < len(nums); i++ {
			// 没有确定最开始选谁，就先for一次看
			chosed = append(chosed, nums[i])
			tmp += nums[i]
			// 从后面的开始选
			dfs(i+1, chosed, tmp)
			if len(chosed) > 0 {
				chosed = chosed[:len(chosed)-1]
				tmp -= nums[i]
			}
		}
	}
	dfs(0, []int{}, 0)
	return res
}
