package leetcode

/*
 * @lc app=leetcode id=129 lang=golang
 *
 * [129] Sum Root to Leaf Numbers
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

func sumNumbers(root *TreeNode) int {
	// return dfs(root, 0)
	return bfs(root)
}

func bfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	sum := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left == nil && node.Right == nil {
			sum += node.Val
		}
		if node.Left != nil {
			node.Left.Val = add(node.Val, node.Left.Val)
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			node.Right.Val = add(node.Val, node.Right.Val)
			queue = append(queue, node.Right)
		}
	}
	return sum
}

func dfs(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return add(sum, root.Val)
	}

	leftSum := dfs(root.Left, add(sum, root.Val))
	rightSum := dfs(root.Right, add(sum, root.Val))
	return leftSum + rightSum
}

func add(sum, val int) int {
	return sum*10 + val
}

// @lc code=end
