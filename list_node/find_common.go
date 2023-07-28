package listnode

func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	p1 := pHead1
	p2 := pHead2
	len1 := getLen(p1)
	len2 := getLen(p2)
	if len1 > len2 {
		for i := 0; i < len1-len2; i++ {
			p1 = p1.Next
		}
	} else if len2 > len1 {
		for i := 0; i < len2-len1; i++ {
			p2 = p2.Next
		}
	}
	return findCommonWithSameBeginning(p1, p2)
}

func getLen(node *ListNode) int {
	p := node
	res := 0
	for p != nil {
		res++
		p = p.Next
	}
	return res
}
func findCommonWithSameBeginning(p1, p2 *ListNode) *ListNode {
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			p1 = p1.Next
			p2 = p2.Next
		} else {
			res := p1
			for p1 != nil && p2 != nil {
				if p1.Val == p2.Val {
					p1 = p1.Next
					p2 = p2.Next
				}
			}
			if p1 == nil && p2 == nil {
				return res
			}
		}
	}
	return nil
}
