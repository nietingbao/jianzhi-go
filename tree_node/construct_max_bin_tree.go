package treenode

func constructMaximumBinaryTree(nums []int) *TreeNode {
	// 递归
	// 给定数组，首先找到最大值所在的位置
	// 生成根节点
	// 树的左边为nums[:index-1],注意边界
	// 树的右边为nums[index+1:],注意边界
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{
			Val: nums[0],
		}
	}
	max := nums[0]
	indexOfMax := 0
	for i, num := range nums {
		if num > max {
			max = num
			indexOfMax = i
		}
	}
	root := &TreeNode{Val: max}
	if indexOfMax != 0 {
		root.Left = constructMaximumBinaryTree(nums[:indexOfMax])
	}
	if indexOfMax != len(nums)-1 {
		root.Right = constructMaximumBinaryTree(nums[indexOfMax+1:])
	}
	return root
}

func indPath(root *TreeNode, target int) [][]int {
	// 二叉树 找路径为目标的所有路径，从根节点开始
	// 深度优先 + 回溯
	res := [][]int{}
	path := []int{}
	var dfs func(*TreeNode, int, []int)
	dfs = func(root *TreeNode, target int, path []int) {
		if root == nil {
			return
		}
		// 将节点加入到path中
		path = append(path, root.Val)
		target = target - root.Val
		if target == 0 && root.Left == nil && root.Right == nil {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		dfs(root.Left, target, path)
		dfs(root.Right, target, path)
		if len(path) > 0 {
			path = path[:len(path)-1]
		}
	}
	dfs(root, target, path)
	return res
}
