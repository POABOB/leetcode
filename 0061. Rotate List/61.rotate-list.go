package leetcode

/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	fast, slow := head, head
	p := head
	nums := 0
	for p != nil {
		nums++
		p = p.Next
	}
	k %= nums
	if nums == 1 || k == 0 {
		return head
	}
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	newHead := slow.Next
	slow.Next = nil
	fast.Next = head
	return newHead
}

// @lc code=end
