package dynamicprogramming

func jumpFloorII(number int) int {
	// write code here
	res := map[int]int{}
	res[1] = 1
	res[2] = 2
	if number <= 2 {
		return res[number]
	}
	for i := 3; i <= number; i++ {
		for j := 1; j < i; j++ {
			res[i] += res[j]
		}
		res[i] += 1
	}
	return res[number]
}
