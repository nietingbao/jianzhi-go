package treenode

func hasPathSum(root *TreeNode, sum int) bool {
	// write code here
	// 空间 O(n), time O(n) => space O(logn), time O(n)
	if root == nil {
		return false
	}
	if root.Val == sum && root.Left == nil && root.Right == nil {
		return true
	}
	var leftHas, rightHas bool
	if root.Left != nil {
		leftHas = hasPathSum(root.Left, sum-root.Val)
	}
	if root.Right != nil {
		rightHas = hasPathSum(root.Right, sum-root.Val)
	}
	return leftHas || rightHas
}
