package listnode

func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	if pHead1 == nil {
		return pHead2
	}
	if pHead2 == nil {
		return pHead1
	}
	p1, p2 := pHead1, pHead2
	var res, end *ListNode
	for p1 != nil && p2 != nil {
		if p1.Val >= p2.Val {
			nextP2 := p2.Next
			if res == nil {
				res = p2
				end = p2
				res.Next = nil
				end.Next = nil
			} else {
				p2.Next = nil
				end.Next = p2
				end = end.Next
			}
			p2 = nextP2
		} else {
			nextP1 := p1.Next
			if res == nil {
				res = p1
				end = p1
				res.Next = nil
				end.Next = nil
			} else {
				p1.Next = nil
				end.Next = p1
				end = end.Next
			}
			p1 = nextP1
		}
	}
	if p1 == nil {
		end.Next = p2
	}
	if p2 == nil {
		end.Next = p1
	}
	return res
}
