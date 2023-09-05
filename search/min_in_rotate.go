package search

func minNumberInRotateArray(nums []int) int {
	// write code here
	if len(nums) == 1 {
		return nums[0]
	}
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}
	i, j := 0, len(nums)-1
	for j > i {
		mid := (i + j) / 2
		if nums[mid] > nums[i] {
			i = mid
		} else if nums[mid] == nums[i] {
			if mid == i {
				if nums[i] > nums[j] {
					return nums[j]
				} else {
					return nums[i]
				}
			} else {
				min := nums[i]
				for k := i; k <= j; k++ {
					if nums[k] < min {
						min = nums[k]
					}
				}
				return min
			}
		} else {
			j = mid
		}
	}
	return 0
}
