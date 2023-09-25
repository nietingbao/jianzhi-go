package dynamicprogramming

import (
	"fmt"
	"strconv"
)

func solve(nums string) int {
	// write code here
	// 能和i-1凑的话，就是di-2
	// 自己单独的话，就是di-1 + 1
	// di 表示i的结果
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums == "0" {
			return 0
		}
		return 1
	}
	d := map[int]int{}
	d[-1] = 1
	d[0] = 1
	for i := 1; i < len(nums); i++ {
		current := fmt.Sprintf("%c", nums[i])
		before := fmt.Sprintf("%c", nums[i-1])
		cNum, err := strconv.ParseInt(current, 10, 64)
		if err != nil {
			fmt.Printf("convert char faild %s", current)
			return 0
		}
		bNum, err := strconv.ParseInt(before, 10, 64)
		if err != nil {
			fmt.Printf("convert char faild %s", before)
			return 0
		}
		d[i] = 0
		if cNum != 0 {
			d[i] += d[i-1]
		}
		sum := bNum*10 + cNum
		if sum == 0 {
			return 0
		}
		if sum <= 26 && bNum != 0 {
			d[i] += d[i-2]
		}
	}
	return d[len(nums)-1]
}
