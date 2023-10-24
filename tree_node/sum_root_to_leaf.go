package treenode

func sumRootToLeaf(root *TreeNode) int {
	// dfs
	// 由于是得到和，所以，遍历的过程中把这个和计算一下即可
	var res int
	if root == nil {
		return 0
	}
	// dfs传入上面的prefix，内部，如果是叶节点，直接加和，否则继续传prefix
	var dfs func(root *TreeNode, prefix int)
	dfs = func(root *TreeNode, prefix int) {
		if root == nil {
			return
		}
		current := prefix<<1 | root.Val
		if isLeaf(root) {
			res += current
			return
		} else {
			dfs(root.Left, current)
			dfs(root.Right, current)
		}
	}
	dfs(root, 0)
	return res
}
