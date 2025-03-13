package leetcode

/*
 * @lc app=leetcode id=138 lang=golang
 *
 * [138] Copy List with Random Pointer
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	hashmap := make(map[*Node]*Node) // map[oldNode]newNode
	dummy := head
	for head != nil {
		hashmap[head] = &Node{Val: head.Val}
		head = head.Next
	}

	head = dummy
	for head != nil {
		hashmap[head].Next = hashmap[head.Next]
		hashmap[head].Random = hashmap[head.Random]
		head = head.Next
	}
	return hashmap[dummy]
}

// @lc code=end
