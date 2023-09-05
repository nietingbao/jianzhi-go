package stackpq

func maxInWindows(num []int, size int) []int {
	// write code here
	if size > len(num) {
		return []int{}
	}
	if size == 0 {
		return []int{}
	}
	var max int
	// keep current window
	window := []int{}
	for i := 0; i < size; i++ {
		if num[i] >= max {
			max = num[i]
		}
		window = append(window, num[i])
	}
	// keep result
	res := []int{}
	res = append(res, max)
	for i := size; i < len(num); i++ {
		head := window[0]
		window = window[1:]
		window = append(window, num[i])
		if num[i] >= max {
			max = num[i]
		} else {
			if head == max {
				tmpMax := window[0]
				for j := 1; j < size; j++ {
					if window[j] > tmpMax {
						tmpMax = window[j]
					}
				}
				max = tmpMax
			}
		}
		res = append(res, max)

	}
	return res
}
