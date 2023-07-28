package treenode

import "fmt"

func Print(pRoot *TreeNode) [][]int {
	// write code here
	res := [][]int{}
	if pRoot == nil {
		return res
	}
	// print node and put left node first
	stackL := make([]*TreeNode, 0)
	// print node and put right node first
	stackR := make([]*TreeNode, 0)
	stackL = append(stackL, pRoot)
	for len(stackR) != 0 || len(stackL) != 0 {
		tmp := []int{}
		for len(stackR) != 0 {
			n := stackR[len(stackR)-1]
			tmp = append(tmp, n.Val)
			if n.Right != nil {
				stackL = append(stackL, n.Right)
			}
			if n.Left != nil {
				stackL = append(stackL, n.Left)
			}

			stackR = stackR[:len(stackR)-1]

			fmt.Print(len(stackR))
		}
		if len(tmp) != 0 {
			res = append(res, tmp)
			continue
		}
		for len(stackL) != 0 {
			n := stackL[len(stackL)-1]
			tmp = append(tmp, n.Val)
			if n.Left != nil {
				stackR = append(stackR, n.Left)
			}
			if n.Right != nil {
				stackR = append(stackR, n.Right)
			}

			stackL = stackL[:len(stackL)-1]

			fmt.Print(len(stackL))
		}
		res = append(res, tmp)
	}
	return res
}
