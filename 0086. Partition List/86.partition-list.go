package leetcode

/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
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

func partition(head *ListNode, x int) *ListNode {
	dummy1, dummy2 := &ListNode{}, &ListNode{}
	p1, p2 := dummy1, dummy2
	for head != nil {
		if head.Val >= x {
			p2.Next = head
			p2 = p2.Next
		} else {
			p1.Next = head
			p1 = p1.Next
		}
		head = head.Next
	}
	p2.Next = nil
	p1.Next = dummy2.Next
	return dummy1.Next
}

// @lc code=end
