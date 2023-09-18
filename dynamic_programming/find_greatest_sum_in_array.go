package dynamicprogramming

func FindGreatestSumOfSubArray(array []int) int {
	// write code here
	// d[i] 到i结束的最大连续子数组的和，max记录之前的最大和
	// 妈的想复杂了
	// di = max(d[i-1]+ array[i], array[i])
	l := len(array)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return array[0]
	}
	d := map[int]int{}
	d[0] = array[0]
	max := d[0]
	for i := 1; i < l; i++ {
		if d[i-1]+array[i] > array[i] {
			d[i] = d[i-1] + array[i]
		} else {
			d[i] = array[i]
		}
		if d[i] > max {
			max = d[i]
		}
	}
	return max
}
