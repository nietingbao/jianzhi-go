package sort_algo

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	SortQ(a)
	fmt.Println(a)
	b := []int{1, 3, 4, 2, 5, 6, 7}
	SortQ(b)
	fmt.Println(b)
	c := []int{7, 6, 5, 4, 3}
	SortQ(c)
	fmt.Println(c)
}

func TestMergeSort(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	MergeSort(a)
	fmt.Println(a)
	b := []int{1, 3, 4, 2, 5, 6, 7}
	MergeSort(b)
	fmt.Println(b)
	c := []int{7, 6, 5, 4, 3}
	MergeSort(c)
	fmt.Println(c)
}
