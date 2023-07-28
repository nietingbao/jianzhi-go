package treenode

func PrintFromTopToBottom(root *TreeNode) []int {
	// write code here
	if root == nil {
		return []int{}
	}
	var pq []*TreeNode
	pq = append(pq, root)
	res := []int{}
	for len(pq) > 0 {
		node := pq[0]
		res = append(res, node.Val)
		if node.Left != nil {
			pq = append(pq, node.Left)
		}
		if node.Right != nil {
			pq = append(pq, node.Right)
		}
		if len(pq) == 1 {
			break
		}
		pq = pq[1:]
	}
	return res
}
