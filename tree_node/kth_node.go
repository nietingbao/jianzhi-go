package treenode

func KthNode(proot *TreeNode, k int) int {
	// write code here
	// 二叉搜索树，左小右大，深度优先遍历，最左遍历完之后，遍历数+1，按中序遍历的方式即可
	if k <= 0 {
		return -1
	}
	var res *TreeNode
	var dfs func(root *TreeNode, num int)
	dfs = func(root *TreeNode, num int) {
		if root == nil {
			return
		}
		// 出口
		if num == k-1 {
			res = root
			return
		}
		dfs(root.Left, num)
		num++
		dfs(root.Right, num)
	}
	dfs(proot, 0)
	if res != nil {
		return res.Val
	}
	return -1
}
