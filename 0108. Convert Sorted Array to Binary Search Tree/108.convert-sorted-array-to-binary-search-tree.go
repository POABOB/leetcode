package leetcode

/*
@lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
*/

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return process(&nums, 0, len(nums)-1)
}

func process(nums *[]int, start int, end int) *TreeNode {
	mid := start + (end-start)>>1
	root := &TreeNode{Val: (*nums)[mid]}
	root.Left = process(nums, start, mid-1)
	root.Right = process(nums, mid+1, end)
	return root
}

// @lc code=end
