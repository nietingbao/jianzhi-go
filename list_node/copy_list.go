package listnode

type RandomListNode struct {
	Label  int
	Next   *RandomListNode
	Random *RandomListNode
}

/**
 *
 * @param pHead RandomListNode类
 * @return RandomListNode类
 */
func Clone(head *RandomListNode) *RandomListNode {
	//write your code here
	res := cloneNext(head)
	cloneRandom(head, res)
	return res
}

func cloneNext(head *RandomListNode) *RandomListNode {
	p := head
	var res, pr *RandomListNode
	for p != nil {
		n := &RandomListNode{Label: p.Label}
		if res == nil {
			res, pr = n, n
		} else {
			pr.Next = n
			pr = pr.Next
		}
		p = p.Next
	}
	return res
}

func stepToGetRandom(head, node *RandomListNode) int {
	if node.Random == nil {
		return -1
	}
	p := head
	res := 0
	for p != nil && p != node.Random {
		p = p.Next
		res++
	}
	if p == nil {
		return -1
	}
	return res
}

func findNodeForStep(head *RandomListNode, step int) *RandomListNode {
	if step == -1 {
		return nil
	}
	p := head
	for step > 0 {
		p = p.Next
		step--
	}
	return p
}

func cloneRandom(src, tar *RandomListNode) {
	ps, pt := src, tar
	for ps != nil {
		step := stepToGetRandom(src, ps)
		if step == -1 {
			pt.Random = nil
		} else {
			pt.Random = findNodeForStep(tar, step)
		}
		ps = ps.Next
		pt = pt.Next
	}
}
