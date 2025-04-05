package leetcode

/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
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

// Method 1. Top Down
// func sortList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}

// 	mid := getMid(head)
// 	head2 := mid.Next
// 	mid.Next = nil
// 	return merge(sortList(head), sortList(head2))
// }

// func getMid(root *ListNode) *ListNode {
// 	fast, slow := root, root
// 	for fast.Next != nil && fast.Next.Next != nil {
// 		fast = fast.Next.Next
// 		slow = slow.Next
// 	}
// 	return slow
// }

// func merge(l1, l2 *ListNode) *ListNode {
// 	dummy := &ListNode{}
// 	curr := dummy
// 	for l1 != nil && l2 != nil {
// 		if l1.Val > l2.Val {
// 			curr.Next = l2
// 			l2 = l2.Next
// 		} else {
// 			curr.Next = l1
// 			l1 = l1.Next
// 		}
// 		curr = curr.Next
// 	}

// 	if l1 != nil {
// 		curr.Next = l1
// 	} else if l2 != nil {
// 		curr.Next = l2
// 	}
// 	return dummy.Next
// }

// Method 2. Bottom Up
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	listLen := 0
	curr := head
	for curr != nil {
		listLen++
		curr = curr.Next
	}

	dummy := &ListNode{Next: head}
	for i := 1; i <= listLen; i <<= 1 {
		curr = dummy.Next
		tail := dummy
		for curr != nil {
			left := curr
			right := nextN(left, i)
			curr = nextN(right, i)
			h, t := merge(left, right)
			// 因為 h 已經是斷開後的 sorted list 必須把被個別組別連接
			// E.g. {a1 <-> a2} {a3 -> a4 -> ... -> an-1 -> an}		  	=> 一開始各組為獨立的
			//    tail.Next tail
			//      {a1 -> a2} -> {a3 -> a4} {an-1 -> ... -> an}		=> tail.Next = h，產生新的組別後，讓組別連接
			//            tail tail.Next
			//      {a1 -> a2} -> {a3 -> a4} {an-1 -> ... -> an}		=> tail = t，接續往下傳遞
			//                 tail.Next tail
			tail.Next = h
			tail = t
		}

	}
	return dummy.Next
}

func nextN(root *ListNode, n int) *ListNode {
	if root == nil {
		return nil
	}
	for n > 1 && root.Next != nil {
		root = root.Next
		n--
	}
	next := root.Next
	root.Next = nil
	return next
}

func merge(l1, l2 *ListNode) (*ListNode, *ListNode) {
	dummy := &ListNode{}
	curr := dummy
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			curr.Next = l2
			l2 = l2.Next
		} else {
			curr.Next = l1
			l1 = l1.Next
		}
		curr = curr.Next
	}

	if l1 != nil {
		curr.Next = l1
	} else if l2 != nil {
		curr.Next = l2
	}

	// tail
	for curr.Next != nil {
		curr = curr.Next
	}
	return dummy.Next, curr
}

// @lc code=end
