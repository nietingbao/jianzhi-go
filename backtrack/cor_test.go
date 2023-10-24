package backtrack

import (
	"fmt"
	"testing"
)

func TestCor(t *testing.T) {
	s := "(0123)"
	res := ambiguousCoordinates(s)
	fmt.Println(res)
	s = "(123)"
	res = ambiguousCoordinates(s)
	fmt.Println(res)
}
