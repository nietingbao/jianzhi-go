package treenode

func longestZigZag(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 0
	}
	r := 0
	dfs(root, -1, 0, &r)
	dfs(root, 1, 0, &r)
	return r
}

// 深度优先遍历二叉树,dir表示从root开始向哪个方向走
func dfs(root *TreeNode, dir int, deep int, r *int) {
	if root == nil {
		return
	}
	if deep > *r {
		*r = deep
	}
	if dir == 1 {
		dfs(root.Right, -1, deep+1, r)
		dfs(root.Right, 1, 0, r)
	} else {
		dfs(root.Left, 1, deep+1, r)
		dfs(root.Left, -1, 0, r)
	}
}

func pathTarget(root *TreeNode, target int) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	tmp := make([]int, 0)
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, target int) {
		if root == nil {
			return
		}
		tmp = append(tmp, root.Val)
		target -= root.Val
		if target == 0 && root.Left == nil && root.Right == nil {
			res = append(res, append([]int(nil), tmp...))
		}

		dfs(root.Left, target)
		dfs(root.Right, target)
		if len(tmp) >= 1 {
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(root, target)
	return res
}
