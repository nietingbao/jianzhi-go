package backtrack

import (
	"fmt"
	"testing"
)

func TestGetNums(t *testing.T) {
	nums := []int{1, 2, 4, 8}
	res := getNums(2, nums, 12)
	for _, re := range res {
		fmt.Println(re)
	}
}
