package listnode

func FindKthToTail(pHead *ListNode, k int) *ListNode {
	// write code here
	len := GetLen(pHead)
	if k > len {
		return nil
	}
	p := pHead
	for i := 0; i < len-k; i++ {
		p = p.Next
	}
	return p
}

func GetLen(pHead *ListNode) int {
	p := pHead
	len := 0
	for p != nil {
		len++
		p = p.Next
	}
	return len
}
