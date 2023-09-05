package search

import "fmt"

func GetNumberOfK(nums []int, k int) int {
	// write code here
	if len(nums) == 0 {
		return 0
	}
	if k < nums[0] || k > nums[len(nums)-1] {
		return 0
	}
	findKIndex := -1
	i, j := 0, len(nums)-1
	for j >= i {
		fmt.Println(i, j)
		if i == j {
			if nums[i] == k {
				findKIndex = i
			} else {
				break
			}
		}
		if nums[i] == k {
			findKIndex = i
			break
		}
		if nums[j] == k {
			findKIndex = j
			break
		}
		if nums[j] < k || nums[i] > k {
			break
		}
		mid := (i + j) / 2
		if nums[mid] == k {
			findKIndex = k
			break
		} else if nums[mid] > k {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	if findKIndex == -1 {
		return 0
	}
	l, r := findKIndex, findKIndex
	res := 0
	for l >= 0 && nums[l] == k {
		l--
		res++
	}
	for r <= len(nums)-1 && nums[r] == k {
		r++
		res++
	}
	res--
	return res
}
