package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result
	carrier := 0

	for l1 != nil || l2 != nil || carrier != 0 {
		if l1 != nil {
			carrier += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carrier += l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: carrier % 10}
		carrier /= 10
		current = current.Next
	}
	return result.Next
}
