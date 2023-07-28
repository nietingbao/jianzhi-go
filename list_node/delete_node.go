package listnode

func deleteNode(head *ListNode, val int) *ListNode {
	// write code here
	// 注意头中尾的不同处理方法
	var res, tail *ListNode
	p := head
	for p != nil {
		if p.Val == val {
			// 开头就是这个
			if res == nil {
				return p.Next
			} else { // 中间找到了
				tail.Next = p.Next
				p.Next = nil
				break
			}
		} else { //正常遍历
			if res == nil {
				res = p
			}
			if tail == nil {
				tail = p
			} else {
				tail = tail.Next
			}
			p = tail.Next
		}
	}
	return res
}
