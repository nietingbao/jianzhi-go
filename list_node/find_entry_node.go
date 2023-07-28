package listnode

func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	p := pHead
	if p.Next == nil {
		return nil
	}
	for p.Next != nil {
		if p.Next == p {
			return p
		}
		n := p.Next
		p.Next = p
		p = n
	}
	return nil
}
