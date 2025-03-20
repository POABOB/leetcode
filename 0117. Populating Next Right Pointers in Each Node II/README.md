# Intuition

給定一棵 `二元樹`，其中每個節點 `包含一個 Next 指標`，請將每個節點的 Next `指向其右側的相鄰節點`。如果沒有右側相鄰節點，則設為 nil。
- 注意：
    - 這棵樹並 `不是完全二元樹`。
    - 必須使用 `O(1) 額外空間` 來完成。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

- 這題可以使用 BFS (廣度優先搜尋, 層序遍歷) 或 DFS (深度優先搜尋, 遞迴) 來解決。

## BFS (廣度優先搜尋, 層序遍歷)

- 使用 `Queue` 來存儲當前層的所有節點。
- 遍歷當前層的節點，將 `Next 指向其右側的相鄰節點`。
    - 依序將 `左右子節` 點加入 Queue，繼續處理下一層。
    - 直到 Queue 為空時結束。

## DFS (深度優先搜尋, 遞迴)

- 使用 Hashmap 記錄每一層最右側的節點。
- 當前節點的 Next `指向該層最後一個記錄的節點`。
- 依序遞迴處理 `左子節點` 和 `右子節點`，以確保 `從左到右遍歷節點`。

<!-- Describe your approach to solving the problem. -->

# Complexity

- Time complexity
    - Method 1. O(n)
    - Method 2. O(n)

<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity
    - Method 1. O(n)
    - Method 2. O(n)

<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code

```go
package leetcode

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
```

```java

import java.util.LinkedList;
import java.util.Map;
import java.util.Queue;

class Node {
    public int val;
    public Node left;
    public Node right;
    public Node next;

    public Node() {
    }

    public Node(int _val) {
        val = _val;
    }

    public Node(int _val, Node _left, Node _right, Node _next) {
        val = _val;
        left = _left;
        right = _right;
        next = _next;
    }
};

class Solution {
    public Node connect(Node root) {
        if (root == null) {
            return null;
        }

        // Method 1. DFS
        // connectNextNode(root, new HashMap<>(), 0);
        // return root;


        // Method 2. BFS
        Queue<Node> queue = new LinkedList<>();
        queue.add(root);

        while (!queue.isEmpty()) {
            int size = queue.size();
            for (int i = 0; i < size; i++) {
                Node node = queue.poll();
                if (i < size - 1) {
                    node.next = queue.peek();
                }
                if (node.left != null) queue.add(node.left);
                if (node.right != null) queue.add(node.right);
            }
        }
        return root;
    }

    private void connectNextNode(Node root, Map<Integer, Node> hashmap, int level) {
        if (root == null) {
            return;
        }

        if (!hashmap.containsKey(level)) {
            hashmap.put(level, root);
        } else {
            hashmap.get(level).next = root;
            hashmap.put(level, hashmap.get(level).next);
        }

        connectNextNode(root.left, hashmap, level + 1);
        connectNextNode(root.right, hashmap, level + 1);
    }
}
```