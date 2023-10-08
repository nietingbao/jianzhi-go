package bitcalc

func NumberOf1(n int) int {
	// write code here
	nums := n2b(n)
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			res++
		}
	}
	return res
}

func n2b(n int) []int {
	res := make([]int, 0, 32)
	if n > 0 {
		for n > 0 {
			res = append(res, n%2)
			n = n / 2
		}
		return res
	} else if n == 0 {
		return res
	} else {
		tmp := n2b(0 - n)
		add := 1
		for i := 0; i < 32; i++ {
			tmp[i] = tmp[i] ^ 1
		}
		for i := 0; i < 32 && add == 1; i++ {
			res := tmp[i] + add
			if res == 2 {
				tmp[i] = 0
				add = 1
			} else {
				tmp[i] = res
				add = 0
			}
		}
		return tmp
	}
}
