package sort_algo

import "fmt"

func InversePairs(nums []int) int {
	// write code here
	l := len(nums)
	if l <= 1 {
		return 0
	}
	res := 0
	mergeFind(nums, &res, 0, l-1)
	return res % 1000000007
}

func mergeFind(nums []int, res *int, l, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)/2
	mergeFind(nums, res, l, mid)
	mergeFind(nums, res, mid+1, r)
	merge(nums, res, l, mid, r)
}

func merge(nums []int, res *int, l, mid, r int) {
	if l >= r {
		return
	}
	i := l
	j := mid + 1
	tmp := make([]int, r-l+1)
	k := 0
	for i <= mid && j <= r {
		if nums[i] < nums[j] {
			tmp[k] = nums[i]
			i++
			k++
		} else {
			tmp[k] = nums[j]
			k++
			j++
			fmt.Println("i=, res+=", i, mid-i+1)
			*res += mid - i + 1
		}
	}
	for i <= mid {
		tmp[k] = nums[i]
		k++
		i++
	}
	for j <= r {
		tmp[k] = nums[j]
		k++
		j++
	}
	for h := 0; h <= r-l; h++ {
		nums[l+h] = tmp[h]
	}
	fmt.Println(l, r, mid, *res)
}
