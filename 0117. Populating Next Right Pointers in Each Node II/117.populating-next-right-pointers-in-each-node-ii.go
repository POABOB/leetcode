package leetcode

/*
 * @lc app=leetcode id=117 lang=golang
 *
 * [117] Populating Next Right Pointers in Each Node II
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	// Method 1. DFS
	// hashmap := make(map[int]*Node)
	// connectNextNode(root, hashmap, 0)
	// return root

	// Method 2. BFS
	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		qLen := len(queue)
		for i := 0; i < qLen; i++ {
			node := queue[i]

			if i+1 < qLen {
				node.Next = queue[i+1]
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[qLen:]
	}
	return root
}

// DFS
// func connectNextNode(root *Node, hashmap map[int]*Node, level int) {
// 	if root == nil {
// 		return
// 	}

// 	if _, existed := hashmap[level]; !existed {
// 		hashmap[level] = root
// 	} else {
// 		hashmap[level].Next = root
// 		hashmap[level] = hashmap[level].Next
// 	}

// 	connectNextNode(root.Left, hashmap, level+1)
// 	connectNextNode(root.Right, hashmap, level+1)
// }

// @lc code=end
