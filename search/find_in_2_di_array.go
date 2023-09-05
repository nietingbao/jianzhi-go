package search

func Find(target int, array [][]int) bool {
	// write code here
	if len(array) == 0 || len(array[0]) == 0 {
		return false
	}
	i, j := len(array)-1, 0
	for i >= 0 && j < len(array[0]) {
		if array[i][j] > target {
			i--
		} else if array[i][j] == target {
			return true
		} else {
			j++
		}
	}
	return false
}
