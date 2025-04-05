package leetcode

import "container/heap"

/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
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

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	minHeap := make(MinHeap, 0)
	heap.Init(&minHeap)

	for _, list := range lists {
		if list != nil {
			heap.Push(&minHeap, list)
		}
	}

	dummy := &ListNode{}
	current := dummy
	for minHeap.Len() > 0 {
		element := heap.Pop(&minHeap).(*ListNode)
		current.Next = element
		// 由該 list 找最小值或找其他 list 中找出最小值
		if element.Next != nil {
			heap.Push(&minHeap, element.Next)
		}
		current = current.Next
	}
	return dummy.Next
}

type MinHeap []*ListNode

func (mh MinHeap) Less(i, j int) bool {
	return mh[i].Val < mh[j].Val
}

func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh MinHeap) Len() int {
	return len(mh)
}

func (mh *MinHeap) Push(x any) {
	*mh = append(*mh, x.(*ListNode))
}

func (mh *MinHeap) Pop() any {
	x := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return x
}

// @lc code=end
