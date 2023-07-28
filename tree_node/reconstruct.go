package treenode

func reConstructBinaryTree(preOrder []int, vinOrder []int) *TreeNode {
	// write code here
	if len(preOrder) != len(vinOrder) {
		return nil
	}
	if len(preOrder) == 0 {
		return nil
	}
	var res *TreeNode
	for i, vin := range vinOrder {
		if vin == preOrder[0] {
			if res == nil {
				res = &TreeNode{}
				res.Val = vin
			}
			if i == 0 {
				res.Left = nil
			} else {
				res.Left = reConstructBinaryTree(preOrder[1:i+1], vinOrder[:i])
			}
			if i == len(vinOrder)-1 {
				res.Right = nil
			} else {
				res.Right = reConstructBinaryTree(preOrder[i+1:], vinOrder[i+1:])
			}
		}
	}
	return res
}
