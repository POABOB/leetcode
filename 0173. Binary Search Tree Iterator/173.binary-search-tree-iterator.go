package leetcode

/*
 * @lc app=leetcode id=173 lang=golang
 *
 * [173] Binary Search Tree Iterator
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

type BSTIterator struct {
	inorder   []*TreeNode
	currIndex int
}

func Constructor(root *TreeNode) BSTIterator {
	arr := make([]*TreeNode, 0)
	inorder(root, &arr)
	return BSTIterator{inorder: arr, currIndex: 0}
}

func inorder(root *TreeNode, res *[]*TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left, res)
	*res = append(*res, root)
	inorder(root.Right, res)
}

func (this *BSTIterator) Next() int {
	curr := this.inorder[this.currIndex]
	this.currIndex++
	return curr.Val
}

func (this *BSTIterator) HasNext() bool {
	return this.currIndex < len(this.inorder)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end
