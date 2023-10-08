package bitcalc

func Add(num1 int, num2 int) int {
	// write code here
	add, sum := num1, num2
	for add != 0 {
		tmp := add ^ sum
		add = (add & sum) << 1
		sum = tmp
	}
	return sum
}
