package treenode

import "fmt"

func KthNode(proot *TreeNode, k int) int {
	// write code here
	if k <= 0 {
		return -1
	}
	if proot == nil {
		return -1
	}
	leftNum := NodeNum(proot.Left)
	fmt.Println(leftNum)
	if k <= leftNum {
		return KthNode(proot.Left, k)
	} else if k == leftNum+1 {
		return proot.Val
	} else {
		return KthNode(proot.Right, k-leftNum-1)
	}
}

func NodeNum(proot *TreeNode) int {
	res := 0
	if proot == nil {
		return 0
	}
	if proot.Left != nil {
		res += NodeNum(proot.Left)
	}
	if proot.Right != nil {
		res += NodeNum(proot.Right)
	}
	if proot != nil {
		res += 1
	}
	return res
}
