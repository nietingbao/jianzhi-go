package treenode

func goodNodes(root *TreeNode) int {
	// 二叉树，根节点到该节点，没有任何节点的值比节点值更大
	// dfs
	// 需要记录一下到当前节点时之前的最大值
	// 遍历节点时，需要更新最大值
	res := 0
	max := root.Val
	var dfs func(root *TreeNode, max int)
	dfs = func(root *TreeNode, max int) {
		if root == nil {
			return
		}
		if root.Val >= max {
			max = root.Val
			res++
		}
		dfs(root.Left, max)
		dfs(root.Right, max)
	}
	dfs(root, max)
	return res
}
