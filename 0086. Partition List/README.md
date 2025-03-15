# Intuition

這題的目標是將 `Linked List` 依照 x 進行分區，使得所有 `< x 的節點`都出現在 `>= x 的節點之前`，且 `維持原始相對順序`。
我們可以透過 兩個 Linked List 來實作，分別儲存 `< x` 和 `>= x` 的節點，最後再將它們合併。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

- 初始化兩個 dummy 節點：
    - dummy1 用來 `存儲 < x` 的節點
    - dummy2 用來 `存儲 >= x`  的節點
    - p1 與 p2 分別指向 dummy1 和 dummy2，作為 Linked List 的 `建構指標`
- `遍歷` 原始 Linked List  (head)：
    - 若 head.Val < x，將該節點 `加到 dummy1`
    - 若 head.Val >= x，將該節點 `加到 dummy2`
    - 移動 head 指標至下一個節點
- `合併` 兩個 Linked List ：
    - p1.Next = dummy2.Next `將小於 x 的部分接上大於等於 x 的部分`
    - p2.Next = nil `避免循環` 的問題
    - 返回 dummy1.Next 作為新 Linked List 的起點

<!-- Describe your approach to solving the problem. -->

# Complexity

- Time complexity
    - O(n)，只需要遍歷鏈結串列一次。

<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity
    - O(1)，只使用了 4 個的指標變數，沒有額外的記憶體分配。

<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code

```go
package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	dummy1, dummy2 := &ListNode{}, &ListNode{}
	p1, p2 := dummy1, dummy2
	for head != nil {
		if head.Val >= x {
			p2.Next = head
			p2 = p2.Next
		} else {
			p1.Next = head
			p1 = p1.Next
		}
		head = head.Next
	}
	p2.Next = nil
	p1.Next = dummy2.Next
	return dummy1.Next
}
```

```java
public class ListNode {
    int val;
    ListNode next;

    ListNode() {
    }

    ListNode(int val) {
        this.val = val;
    }

    ListNode(int val, ListNode next) {
        this.val = val;
        this.next = next;
    }
}

class Solution {
    public ListNode partition(ListNode head, int x) {
        ListNode dummy1 = new ListNode(), dummy2 = new ListNode();
        ListNode p1 = dummy1, p2 = dummy2;
        while (head != null) {
            if (head.val >= x) {
                p2.next = head;
                p2 = p2.next;
            } else {
                p1.next = head;
                p1 = p1.next;
            }
            head = head.next;
        }
        p2.next = null;
        p1.next = dummy2.next;
        return dummy1.next;
    }
}
```




