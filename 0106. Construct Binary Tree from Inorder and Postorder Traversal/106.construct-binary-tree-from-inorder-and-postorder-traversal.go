package leetcode

/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
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

func buildTree(inorder []int, postorder []int) *TreeNode {
	// 建立 hashmap 儲存 inorder value to key
	hashmap := make(map[int]int, len(inorder))
	for k, v := range inorder {
		hashmap[v] = k
	}
	return build(&inorder, &postorder, hashmap, 0, len(inorder)-1)
}

func build(inorder *[]int, postorder *[]int, hashmap map[int]int, inStart, inEnd int) *TreeNode {
	if inStart > inEnd {
		return nil
	}

	// postorder 要從元素最後一位取
	postLen := len((*postorder))
	rootValue := (*postorder)[postLen-1]
	(*postorder) = (*postorder)[:postLen-1]
	index := hashmap[rootValue]
	root := &TreeNode{Val: rootValue}
	// 這邊會先建構右子樹，原因是 (*postorder)[postLen-1] 每次都取最後一位
	// postorder 內 [{左},{右},root] 的結構
	// 如果由後面往前取的話，那麼每次找到的跟都會先是右子樹的根
	root.Right = build(inorder, postorder, hashmap, index+1, inEnd)
	root.Left = build(inorder, postorder, hashmap, inStart, index-1)
	return root
}

// @lc code=end
