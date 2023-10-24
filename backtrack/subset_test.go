package backtrack

import (
	"fmt"
	"testing"
)

func TestSubset(t *testing.T) {
	nums := []int{1, 2, 3}
	res := subsets(nums)
	fmt.Println(res)
	nums = []int{1, 2, 3}
	res = subsetsWithDup(nums)
	fmt.Println(res)
}
