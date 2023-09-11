package stackpq

func IsPopOrder(pushV []int, popV []int) bool {
	// write code here
	var stack []int
	if len(pushV) != len(popV) {
		return false
	}
	if len(popV) == 0 {
		return false
	}
	i := 0
	for j := 0; j < len(pushV); j++ {
		if pushV[j] != popV[i] {
			stack = append(stack, pushV[j])
		} else {
			i++
			for {
				if len(stack) > 0 {
					if stack[len(stack)-1] == popV[i] {
						stack = stack[:len(stack)-1]
						i++
						if i == len(popV) {
							return true
						}
					} else {
						break
					}
				} else {
					break
				}
			}

		}
	}
	for j := len(stack) - 1; j >= 0; j-- {
		if stack[j] != popV[i] {
			return false
		}
		i++
	}
	return true
}
