package dynamicprogramming

func FindGreatestSumArrOfSubArray(array []int) []int {
	// write code here
	// d[i]记录包括i的最大连续子数组的和，max表示过去的最大和
	// left，right表示最大连续子数组的
	// di = max(di-1 + ai, ai)
	// 何时更新left？ai==d[i]
	// 何时更新right？left更新或者，d[i] ==d[i-1] + ai,同时max更大或者保持一致？
	l := len(array)
	if l == 0 {
		return []int{}
	}
	if l == 1 {
		return array
	}
	d := map[int]int{}
	d[0] = array[0]
	max := d[0]
	left, right := 0, 0
	for i := 1; i < l; i++ {
		// 添加进来了
		if d[i-1]+array[i] > array[i] {
			d[i] = d[i-1] + array[i]
			// 更新了max
			if d[i] >= max {
				right = i
				max = d[i]
			}
		} else if d[i-1]+array[i] == array[i] {
			d[i] = array[i]
			if d[i] == max {
				right = i
			}
		} else {
			// 过去的不要了，直接从新起点开始
			d[i] = array[i]
			if d[i] > max {
				max = d[i]

			}
			if array[i] < d[i-1] {
				continue
			}
			left = i
			right = i
		}
	}
	return array[left : right+1]
}
