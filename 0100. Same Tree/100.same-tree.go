package leetcode

/*
 * @lc app=leetcode id=100 lang=golang
 *
 * [100] Same Tree
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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Method 1. DFS
	// if p == nil && q == nil {
	// 	return true
	// }
	// if p == nil || q == nil || p.Val != q.Val {
	// 	return false
	// }
	// return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)

	// Method 2. BFS
	queue1, queue2 := make([]*TreeNode, 0), make([]*TreeNode, 0)
	queue1 = append(queue1, p)
	queue2 = append(queue2, q)
	for len(queue1) > 0 {
		node1 := queue1[0]
		node2 := queue2[0]
		queue1 = queue1[1:]
		queue2 = queue2[1:]
		if node1 == nil && node2 == nil {
			continue
		}
		if node1 == nil || node2 == nil || node1.Val != node2.Val {
			return false
		}
		queue1 = append(queue1, []*TreeNode{node1.Left, node1.Right}...)
		queue2 = append(queue2, []*TreeNode{node2.Left, node2.Right}...)
	}
	return true
}

// @lc code=end
