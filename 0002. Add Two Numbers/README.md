# Intuition

這題是對兩個 `逆序` 的 `非負整數` Linked List 進行相加，並 `返回其總和的 Linked List`。
由於 Linked List 的儲存順序是 `低位在前`，`高位在後`，我們可以從頭部開始依次相加，模擬小學算術的 `逐位相加` 過程。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

- 使用 `虛擬頭節點 (dummy node)` 來簡化結果 Linked List 的構建。
- 設定 `carry` 變數來 `處理進位`，並逐位相加 l1 和 l2 的值。
- 若 l1 或 l2 `任一為 nil`，則視其值為 `0`。
- 在 l1 和 l2 遍歷完成後，若 `carry > 0`，則 `新增一個節點存儲 carry`。

<!-- Describe your approach to solving the problem. -->

# Complexity

- Time complexity
  - O(Max(m,n))，其中 m 和 n 分別是 l1 和 l2 的長度，我們最多遍歷較長的 Linked List 一次。
- Space complexity
  - O(Max(m,n))，我們建立了一個新的 Linked List 來存儲結果。

<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code

```go
package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += carry
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}
	if carry > 0 {
		current.Next = &ListNode{Val: carry}
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
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode dummy = new ListNode(0);
        ListNode current = dummy;
        int carry = 0;
        while (l1 != null || l2 != null) {
            int sum = carry;
            if (l1 != null) {
                sum += l1.val;
                l1 = l1.next;
            }
            if (l2 != null) {
                sum += l2.val;
                l2 = l2.next;
            }
            carry = sum / 10;
            current.next = new ListNode(sum % 10);
            current = current.next;
        }
        if (carry > 0) {
            current.next = new ListNode(carry);
        }
        return dummy.next;
    }
}
```