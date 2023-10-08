package sort_algo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInverse(t *testing.T) {
	nums := []int{2, 7, 4, 8, 3, 5, 6}
	cal := InversePairs(nums)
	expect := 8
	require.Equal(t, cal, expect)
}
