# Intuition
本題要求 `複製` 一個帶有`隨機指標 (Random Pointer)` 的 `Linked List`，且 新的 Linked List 必須與原 Linked List 具有 `相同的結構與值`。這意味著我們不僅要處理 Next 指標，還需要正確複製 Random 指標。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 這裡使用 Hashmap 來儲存 `舊節點` 與 `新節點` 之間的對應關係。
- 第一遍遍歷：
    - `建立所有新節點`，並存入 `map[舊節點] = 新節點`。
- 第二遍遍歷：
    - 設置 `新節點的 Next 和 Random 指標`，確保結構與原 Linked List 一致。
- 最後返回複製後的 Linked List 的頭部。

<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity
    - O(n)，遍歷了 Linked List 兩次。
- Space complexity 
    - O(n)，使用了 Hashmap 來存儲對應關係。
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```go
package leetcode

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
```

```java
import java.util.HashMap;
import java.util.Map;

class Node {
    int val;
    Node next;
    Node random;

    public Node(int val) {
        this.val = val;
        this.next = null;
        this.random = null;
    }
}

class Solution {
    public Node copyRandomList(Node head) {
        Map<Node, Node> hashmap = new HashMap<>();
        Node dummy = head;
        while (head != null) {
            hashmap.put(head, new Node(head.val));
            head = head.next;
        }

        head = dummy;
        while (head != null) {
            hashmap.get(head).next = hashmap.get(head.next);
            hashmap.get(head).random = hashmap.get(head.random);
            head = head.next;
        }
        return hashmap.get(dummy);
    }
}
```