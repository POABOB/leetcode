package leetcode

/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
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

func buildTree(preorder []int, inorder []int) *TreeNode {
	hashmap := make(map[int]int, len(inorder))
	for k, v := range inorder {
		hashmap[v] = k
	}
	return build(&preorder, &inorder, hashmap, 0, len(inorder)-1)
}

func build(preorder *[]int, inorder *[]int, hashmap map[int]int, inStart, inEnd int) *TreeNode {
	if inStart > inEnd {
		return nil
	}
	// preorder 的 root 統一在首位
	rootValue := (*preorder)[0]
	(*preorder) = (*preorder)[1:]
	// 找出 inorder 內 root 的 index
	i := hashmap[rootValue]
	root := &TreeNode{Val: rootValue}
	root.Left = build(preorder, inorder, hashmap, inStart, i-1)
	root.Right = build(preorder, inorder, hashmap, i+1, inEnd)
	return root
}

// @lc code=end
