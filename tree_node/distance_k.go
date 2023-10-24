package treenode

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	// 二叉树，指定节点，节点一定在树里面，距离k的多个节点的值，可能为空
	// 首先先找到target
	// 从target开始dfs
	// dfs先序遍历
	// 如果是要从上往下，就不能先找node，直接从头开始dfs，记录到头结点的距离
	parent := make(map[*TreeNode]*TreeNode)
	var findParent func(root *TreeNode)
	findParent = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			parent[root.Left] = root
			findParent(root.Left)
		}
		if root.Right != nil {
			parent[root.Right] = root
			findParent(root.Right)
		}
	}
	findParent(root)
	parent[root] = nil
	res := []int{}
	var node *TreeNode
	var findTarget func(root, target *TreeNode)
	findTarget = func(root, target *TreeNode) {
		if root == nil {
			return
		}
		if root.Val == target.Val {
			node = root
			return
		}
		findTarget(root.Left, target)
		findTarget(root.Right, target)
	}
	findTarget(root, target)
	if node == nil {
		return res
	}
	// 如何避免循环？光使用方向不对，因为往上到达根部后还可以往下。
	// 记录调用自己的那一方，如果即将调用的不是调用自己的那一方，就可以调用。
	var dfs func(from, root *TreeNode, dis int)
	dfs = func(from, root *TreeNode, dis int) {
		if root == nil {
			return
		}
		if dis == k {
			res = append(res, root.Val)
			return
		}
		// 前序遍历
		// 往下为1，往上为-1
		if root.Left != from {
			dfs(root, root.Left, dis+1)
		}
		if root.Right != from {
			dfs(root, root.Right, dis+1)
		}

		pa := parent[root]
		if pa != from {
			dfs(root, pa, dis+1)
		}
	}
	dfs(nil, node, 0)
	return res
}
