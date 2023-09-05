package treenode

import "fmt"

func lowestCommonAncestor(root *TreeNode, o1 int, o2 int) int {
	// write code here
	// must implement
	parent := map[int]int{}
	if root == nil {
		return 0
	}
	pq := []*TreeNode{}
	pq = append(pq, root)
	parent[root.Val] = 1<<63 - 1
	for {
		_, ok := parent[o1]
		_, ok2 := parent[o2]
		if ok && ok2 {
			break
		}
		fmt.Println("len is, current node is,", len(pq), pq[0].Val)
		node := pq[0]
		if len(pq) == 1 {
			pq = []*TreeNode{}
		} else {
			pq = pq[1:]
		}
		if node.Left != nil {
			parent[node.Left.Val] = node.Val
			pq = append(pq, node.Left)
		}
		if node.Right != nil {
			parent[node.Right.Val] = node.Val
			pq = append(pq, node.Right)
		}
	}
	o1Path := map[int]struct{}{}
	o1Path[o1] = struct{}{}
	for {
		pa, ok := parent[o1]
		if ok {
			o1Path[pa] = struct{}{}
			o1 = pa
		} else {
			break
		}
	}

	for {
		_, ok := o1Path[o2]
		if ok {
			return o2
		} else {
			pa, ok := parent[o2]
			if ok {
				o2 = pa
			} else {
				return 0
			}
		}
	}
}
