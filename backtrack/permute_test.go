package backtrack

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	tar := []int{1, 1, 3}
	res := permute([]int{1, 2, 3})
	fmt.Println(res)
	res = permuteUnique(tar)
	fmt.Println(res)
}
