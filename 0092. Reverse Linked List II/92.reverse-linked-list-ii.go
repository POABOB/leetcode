package leetcode

/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
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

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// Method 1. 遍歷
	// if head == nil || head.Next == nil {
	// 	return head
	// }
	// dummy := &ListNode{0, head}
	// p := dummy
	// for i := 0; i < left-1; i++ {
	// 	p = p.Next
	// }

	// var prev *ListNode = nil
	// curr := p.Next
	// for i := 0; i < right-left+1; i++ {
	// 	next := curr.Next
	// 	curr.Next = prev
	// 	prev = curr
	// 	curr = next
	// }
	// // p.Next 是被 reversed 第一個節點
	// // 第一個節點的 Next 其實一開始會是 nil (prev 初始值 nil)，所以要指回去 reversed 最後節點 curr
	// p.Next.Next = curr
	// // p.Next 要指向 reversed 的最後一個節點 prev
	// p.Next = prev
	// return dummy.Next

	// Method 2. 遞迴
	var successor *ListNode
	return reverseBetweenHelper(head, left, right, &successor)
}

func reverseBetweenHelper(head *ListNode, m int, n int, successor **ListNode) *ListNode {
	// base case，找到要反轉的位置
	if m == 1 {
		return reverseN(head, n, successor)
	}
	// 如果 m != 1，head 就會一直前進，直到找到要反轉的範圍
	head.Next = reverseBetweenHelper(head.Next, m-1, n-1, successor)
	return head
}

func reverseN(head *ListNode, n int, successor **ListNode) *ListNode {
	if n == 1 {
		*successor = head.Next
		return head
	}
	last := reverseN(head.Next, n-1, successor)
	// left = 2, right = 4
	// 1 -> 2 -> 3 -> 4 -> 5
	// n = 1, head = 4, successor = 5
	// n = 2, last = 4, head = 3, 4 -> 3 -> 5
	// n = 3, last = 4, head = 2, 4 -> 3 -> 2 -> 5
	head.Next.Next = head
	head.Next = *successor
	return last
}

// @lc code=end
