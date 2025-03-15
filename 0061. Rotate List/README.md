# Intuition

這題要求我們將一個 Linked List 進行 `右移操作`，將列表的 `最後 k 個節點` 移動到 Linked List 的 `最前面`
，並且需要保證 `最後的節點順序不會被改變`。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

- 由於是右移操作，我們可以先計算出 Linked List 的 `長度 nums`，然後對 k 進行 `取餘運算 (k % n)`，`避免 k 超過長度`。
    - 如果 `k == 0` 或者 Linked List `只有一個節點`，那麼 `直接返回原 Linked List`。
- 使用兩個指針 `fast` 和 `slow`，初始化都 `指向 head`。然後讓 fast 指針 `先向前移動 k 步`。
- 接著，`同時移動` fast 和 slow 直到 `fast 指向` Linked List 的 `最後一個節點`。
- 此時，`slow 指向的節點` 就是 `需要進行斷開的節點`，並且將 `其後面的節點 slow.Next` 設為 `新的頭節點`。
- 最後，將 `fast.Next 指向原來的頭部`，形成循環鏈表結構，並且 `斷開與 slow 節點的連接`。

# Complexity

- Time complexity
    - O(n)，其中 n 是鏈表的長度。

<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity
    - O(1)，因為只用了常數額外空間。

<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code

```go
package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	fast, slow := head, head
	p := head
	nums := 0
	for p != nil {
		nums++
		p = p.Next
	}
	k %= nums
	if nums == 1 || k == 0 {
		return head
	}
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	newHead := slow.Next
	slow.Next = nil
	fast.Next = head
	return newHead
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
    public ListNode rotateRight(ListNode head, int k) {
        if (head == null) {
            return null;
        }
        int nums = 0;
        ListNode p = head;
        while (p != null) {
            nums++;
            p = p.next;
        }
        k = k % nums;
        if (k == 0 || nums == 1) {
            return head;
        }

        ListNode fast = head, slow = head;
        for (int i = 0; i < k; i++) {
            fast = fast.next;
        }
        while (fast.next != null) {
            slow = slow.next;
            fast = fast.next;
        }
        ListNode newHead = slow.next;
        slow.next = null;
        fast.next = head;
        return newHead;
    }
}
```