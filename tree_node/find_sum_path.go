package treenode

import "fmt"

func FindPath(root *TreeNode, target int) [][]int {
	return findPath(root, target, []int{})
}

func findPath(root *TreeNode, target int, cur []int) [][]int {
	var res [][]int
	if root == nil {
		return [][]int{}
	}
	if root.Val == target {
		if root.Left == nil && root.Right == nil {
			cur = append(cur, root.Val)
			res = append(res, cur)
		}
	} else if root.Val != target {
		if root.Left == nil && root.Right == nil {
			return [][]int{}
		}
		cur = append(cur, root.Val)
		if root.Left != nil {
			fmt.Println(root.Val)
			fmt.Println("cur is", cur)
			res = append(res, findPath(root.Left, target-root.Val, cur)...)
		}
		if root.Right != nil {
			res = append(res, findPath(root.Right, target-root.Val, cur)...)
		}
	}
	return res
}
