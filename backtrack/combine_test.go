package backtrack

import (
	"fmt"
	"testing"
)

func TestCombine(t *testing.T) {
	res := combinationSum([]int{2, 3, 6, 7}, 7)
	fmt.Println(res)
	res = combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	fmt.Println(res)
}
