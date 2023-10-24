package dynamicprogramming

import (
	"fmt"
	"testing"
)

func TestCoin(t *testing.T) {
	res := coinChange([]int{1, 2, 5}, 11)
	fmt.Println(res)
	res = coinChange([]int{2}, 3)
	fmt.Println(res)
}
