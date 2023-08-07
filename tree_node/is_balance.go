package treenode

func IsBalanced_Solution(pRoot *TreeNode) bool {
	// write code here
	if pRoot == nil {
		return true
	}
	le := treeHight(pRoot.Left)
	ri := treeHight(pRoot.Right)
	if le == ri || le-ri == 1 || ri-le == 1 {
		return IsBalanced_Solution(pRoot.Left) && IsBalanced_Solution(pRoot.Right)
	}
	return false
}

//{1,2,3,4,5,#,6,#,#,7}

func treeHight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var res int
	res += 1
	var le, ri int
	if root.Right != nil {
		le = treeHight(root.Right)
	}
	if root.Left != nil {
		ri = treeHight(root.Left)
	}
	if le >= ri {
		res += le
	} else {
		res += ri
	}
	return res
}
