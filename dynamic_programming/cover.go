package dynamicprogramming

func rectCover(number int) int {
	// write code here
	// f(n) = f(n-1) + f(n-2)
	// f1 = 1
	// f2 = 2
	if number < 3 {
		return number
	}
	a, b := 1, 2
	for i := 3; i <= number; i++ {
		tmp := b
		b = a + b
		a = tmp
	}
	return b
}
