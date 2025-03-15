# Intuition

這題要求刪除排序 Linked List 中 `所有重複的節點`，使得每個元素只要出現過重複，就不應該出現在結果 Linked List 中。
可以遍歷 Linked List ，當發現某個數字有重複時，就把它完全移除，不留下任何一個相同數值的節點。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

- 使用虛擬頭節點（dummy node）
    - 由於 `頭節點本身可能會被刪除`，因此使用一個 dummy 指向 head，方便統一處理。
- 遍歷 Linked List
    - 使用 prev 指針來標記最後一個 `確定不會被刪除的節點`，確保刪除後的 Linked List 正確連接。
    - 當發現 `curr 和 curr.Next 的值相同` 時，進入 loop，`跳過所有該值的節點`。
        - 當 curr 停止時，它指向的是該值的最後一個節點，因此 `prev.Next 直接指向 curr.Next`，`跳過所有該值的節點`。
    - 如果 `curr 沒有重複值`，則 `prev 指向 curr`，保持 Linked List 結構。

<!-- Describe your approach to solving the problem. -->

# Complexity

- Time complexity
    - O(n)，其中 n 是 Linked List 的長度。

<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity
    - O(1)，因為我們沒有使用額外的資料結構，只用了幾個額外指標。

<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code

```go
package leetcode

func deleteDuplicates(head *ListNode) *ListNode {
	curr := head
	dummy := &ListNode{Next: head}
	prev := dummy

	for curr != nil {
		if curr.Next != nil && curr.Val == curr.Next.Val {
			for curr.Next != nil && curr.Val == curr.Next.Val {
				curr = curr.Next
			}
			prev.Next = curr.Next
		} else {
			prev = prev.Next
		}
		curr = curr.Next
	}
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
    public ListNode deleteDuplicates(ListNode head) {
        ListNode curr = head;
        ListNode dummy = new ListNode(-999, head);
        ListNode prev = dummy;
        while (curr != null) {
            if (curr.next != null && curr.val == curr.next.val) {
                while (curr.next != null && curr.val == curr.next.val) {
                    curr = curr.next;
                }
                prev.next = curr.next;
            } else {
                prev = prev.next;
            }
            curr = curr.next;
        }
        return dummy.next;
    }
}
```




