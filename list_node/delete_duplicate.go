package listnode

import "fmt"

func DeleteDuplication(pHead *ListNode) *ListNode {
	// write code here
	// 注意一开始就重复，注意结尾重复
	var res, tail *ListNode
	p := pHead
	for p != nil {
		avai, pos := FindAvailableNode(p)
		fmt.Printf("current %v, avil %v, pos %v", p, avai, pos)
		if avai == nil {
			break
		} else {
			if res == nil {
				res, tail = avai, avai
			} else {
				tail.Next = avai
				tail = tail.Next
			}
		}
		p = pos
	}
	if tail != nil {
		tail.Next = nil
	}
	return res
}

func FindAvailableNode(node *ListNode) (avai, pos *ListNode) {
	if node.Next == nil {
		return node, nil
	}
	p, pn := node, node.Next
	same := false
	for pn != nil {
		if p.Val == pn.Val {
			same = true
			pn = pn.Next
		} else {
			if same {
				p = pn
				pn = pn.Next
				same = false
			} else {
				return p, pn
			}
		}
	}
	if !same {
		return p, nil
	}
	return nil, nil
}
