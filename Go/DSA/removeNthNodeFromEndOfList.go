// Given the head of a linked list, remove the nth node from the end of the list and return its head.

// Example 1:

// Input: head = [1,2,3,4,5], n = 2
// Output: [1,2,3,5]
// Example 2:

// Input: head = [1], n = 1
// Output: []
// Example 3:

// Input: head = [1,2], n = 1
// Output: [1]

// Constraints:

// The number of nodes in the list is sz.
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz

// Follow up: Could you do this in one pass?

package DSA

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func InitializeListNode(val int) *ListNode {
	return &ListNode{
		Val: val,
	}
}

func DisplayList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func (head *ListNode) AddToList(val int) {
	for head != nil {
		if head.Next == nil {
			head.Next = &ListNode{
				Val: val,
			}
			return
		}
		head = head.Next
	}
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head

	for range n {
		fast = fast.Next
	}

	if fast == nil {
		return head.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return head
}
