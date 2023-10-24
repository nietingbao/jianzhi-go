package bitcalc

func NumberOf1(n int) int {
	if n > 0 {
		return numOf1Pos(n)
	} else if n == 0 {
		return 0
	} else {
		max := 1<<31 - 1
		codeNum := max + n + 1
		return numOf1Pos(codeNum) + 1
	}
}

func numOf1Pos(n int) int {
	res := 0
	for n > 0 {
		if n&1 == 1 {
			res++
		}
		n = n / 2
	}
	return res
}
