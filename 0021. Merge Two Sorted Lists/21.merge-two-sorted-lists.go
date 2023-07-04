package leetcode

/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Method 1. 遍歷
	// 建立一個 head 作為返回的答案
	// head := &ListNode{-1, nil}
	// // dummy 是一個暫存指標
	// dummy := head

	// for list1 != nil && list2 != nil {
	// 	if list1.Val > list2.Val {
	// 		dummy.Next = list2
	// 		list2 = list2.Next
	// 	} else {
	// 		dummy.Next = list1
	// 		list1 = list1.Next
	// 	}
	// 	dummy = dummy.Next
	// }

	// if list1 != nil {
	// 	dummy.Next = list1
	// }

	// if list2 != nil {
	// 	dummy.Next = list2
	// }
	// return head.Next

	// Method 2. 遞迴
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val > list2.Val {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}

	list1.Next = mergeTwoLists(list1.Next, list2)
	return list1
}

// @lc code=end
