package dynamicprogramming

func jumpFloor(number int) int {
	// write code here
	// f(n) = f(n- 1) + f(n-2)
	if number == 1 {
		return 1
	}
	if number == 2 {
		return 2
	}
	a, b := 1, 2
	for i := 3; i <= number; i++ {
		tmp := b
		b = a + b
		a = tmp
	}
	return b
}
