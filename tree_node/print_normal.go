package treenode

func Print(pRoot *TreeNode) [][]int {
	// write code here
	var res [][]int
	if pRoot == nil {
		return res
	}
	var pq1, pq2 []*TreeNode
	pq1 = append(pq1, pRoot)
	for len(pq1) > 0 || len(pq2) > 0 {
		tmp := []int{}
		for len(pq1) > 0 {
			add := pq1[0]
			tmp = append(tmp, add.Val)
			if add.Left != nil {
				pq2 = append(pq2, add.Left)
			}
			if add.Right != nil {
				pq2 = append(pq2, add.Right)
			}
			if len(pq1) == 1 {
				pq1 = []*TreeNode{}
				break
			}
			pq1 = pq1[1:]
		}
		if len(tmp) > 0 {
			res = append(res, tmp)
			continue
		}
		for len(pq2) > 0 {
			add := pq2[0]
			tmp = append(tmp, add.Val)
			if add.Left != nil {
				pq1 = append(pq1, add.Left)
			}
			if add.Right != nil {
				pq1 = append(pq1, add.Right)
			}
			if len(pq2) == 1 {
				pq2 = []*TreeNode{}
				break
			}
			pq2 = pq2[1:]
		}
		if len(tmp) > 0 {
			res = append(res, tmp)
		}
	}
	return res
}
