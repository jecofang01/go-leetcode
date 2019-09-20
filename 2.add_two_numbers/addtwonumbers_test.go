package addtwonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeListNode(numbers ...int) *ListNode {
	result := &ListNode{}
	current := result
	for _, number := range numbers {
		current.Next = &ListNode{Val: number}
		current = current.Next
	}
	return result.Next
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1  *ListNode
		l2  *ListNode
		out *ListNode
	}{
		{makeListNode(2, 4, 3), makeListNode(5, 6, 4), makeListNode(7, 0, 8)},
		{makeListNode(2, 4, 3), makeListNode(5), makeListNode(7, 4, 3)},
		{makeListNode(2), makeListNode(5, 6, 4), makeListNode(7, 6, 4)},
		{makeListNode(2, 4, 7), makeListNode(5, 6, 4), makeListNode(7, 0, 2, 1)},
		{makeListNode(2, 4, 7), makeListNode(0), makeListNode(2, 4, 7)},
		{makeListNode(0), makeListNode(5, 6, 4), makeListNode(5, 6, 4)},
	}

	for _, tt := range tests {
		tt := tt
		t.Run("", func(t *testing.T) {
			actual := AddTwoNumbers(tt.l1, tt.l2)
			expected := tt.out
			for actual != nil && expected != nil {
				assert.Equal(t, expected.Val, actual.Val)
				actual, expected = actual.Next, expected.Next
			}
		})
	}
}
