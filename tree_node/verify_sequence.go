package treenode

import "fmt"

func VerifySquenceOfBST(sequence []int) bool {
	// write code here
	if len(sequence) == 0 {
		return false
	}
	if len(sequence) == 1 {
		return true
	}
	littleIndex := -1
	length := len(sequence)
	root := sequence[length-1]
	for i := length - 2; i >= 0; i-- {
		if sequence[i] < root {
			littleIndex = i
			break
		}
	}
	fmt.Println(littleIndex)
	var left, right bool
	if littleIndex == -1 {
		left = true
		right = VerifySquenceOfBST(sequence[0 : length-1])
	} else {
		for i := littleIndex; i >= 0; i-- {
			if sequence[i] >= root {
				return false
			}
		}
		fmt.Println(left)
		fmt.Println(right)
		left = VerifySquenceOfBST(sequence[0 : littleIndex+1])
		rightS := sequence[littleIndex+1 : length-1]
		if len(rightS) == 0 {
			right = true
		} else {
			right = VerifySquenceOfBST(rightS)
		}
		fmt.Println(left)
		fmt.Println(right)
	}
	return left && right
}
