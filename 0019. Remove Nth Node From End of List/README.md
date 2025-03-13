# Intuition
這題要求刪除單向 `Linked List` 中 `倒數第 n 個節點`。關鍵點在於如何 `有效` 找到該節點並刪除它。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

## 雙指針解法（Two-Pointer Technique）

- 使用 Dummy 節點：
	- 由於我們 `可能會刪除 head 節點`，因此使用一個 dummy 節點來統一 `處理邊界條件`。
	- dummy.next = head，這樣即使 head 被刪除，我們仍然能返回 dummy.next。
- 初始化 `fast` 和 `slow `指針：
	- 兩個指針都從 dummy 開始。
	- fast 先前進 `n+1 步`，使得 `fast 和 slow 指針相距 n`。
	- 移動 fast 和 slow 指針，`直到 fast 到達尾部`
	- slow 指針的 `下一個節點即為倒數第 n 個節點`。
- 刪除節點：
	- 將 slow 指向 slow.Next.Next，即 `刪除該節點`。


<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity
    - O(n)，因為最多遍歷一次。
- Space complexity 
    - O(1)
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```go
package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	fast, slow := dummy, dummy
	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
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
	public ListNode removeNthFromEnd(ListNode head, int n) {
		ListNode dummy = new ListNode(0, head);
		ListNode slow = dummy;
		ListNode fast = dummy;
		for (int i = 0; i < n + 1; i++) {
			fast = fast.next;
		}
		while (fast != null) {
			slow = slow.next;
			fast = fast.next;
		}
		slow.next = slow.next.next;
		return dummy.next;
	}
}
```