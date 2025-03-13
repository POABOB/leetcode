# Intuition
這題要 `反轉單向 Linked List` 中指定範圍的節點，並且需要保持其餘部分的結構不變。由於反轉範圍 `left 到 right` 可能在 Linked List 的中間，因此需要小心 `處理邊界條件`。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

## 方法 1：遍歷（Iterative Approach）

- 使用 dummy 節點 來簡化邊界條件處理。
- 先移動指標到 `left-1 位置`，找到 `left 的前一個節點 p`。
- `反轉` left 到 right 之間的節點。
- 重新 `連接 left 的前一個節點 p` 以及 `right 之後的節點`。

## 方法 2：遞迴（Recursive Approach）

- 這種方法利用遞迴反轉前 n 個節點的技巧，來反轉 left 到 right 之間的節點。
- `基本情況 (Base Case)`：
    - 當 `left == 1`，我們可以直接 `反轉前 n 個節點`，這樣就能夠處理 left = 1 的特殊情況。
- `遞迴步驟 (Recursive Step)`：
    - 將 head.Next 設為遞迴調用的結果。
    - 這樣當 left 遞減到 1 時，就會進入 reverseN 方法，進行 n 節點的反轉。
- `反轉前 N 個節點 (reverseN 方法)`：
    - 當 n == 1 時，儲存 `head.Next 為 successor`，因為反轉後的 `最後一個節點` 需要 `與 successor 連接`。
    - 否則繼續遞迴 `反轉 n-1 個節點`。
    - 反轉後的 `head.Next 需要指回 successor`，確保剩下的部分不受影響。


<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity
    - Method1. O(n)
    - Method2. O(n)
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity 
    - Method1. O(1)
    - Method2. O(n)，Function Call Stack
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```go
package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// Method 1. 遍歷
	// if head == nil || head.Next == nil {
	// 	return head
	// }
	// dummy := &ListNode{0, head}
	// p := dummy
	// for i := 0; i < left-1; i++ {
	// 	p = p.Next
	// }

	// var prev *ListNode = nil
	// curr := p.Next
	// for i := 0; i < right-left+1; i++ {
	// 	next := curr.Next
	// 	curr.Next = prev
	// 	prev = curr
	// 	curr = next
	// }
	// // p.Next 是被 reversed 第一個節點
	// // 第一個節點的 Next 其實一開始會是 nil (prev 初始值 nil)，所以要指回去 reversed 最後節點 curr
	// p.Next.Next = curr
	// // p.Next 要指向 reversed 的最後一個節點 prev
	// p.Next = prev
	// return dummy.Next

	// Method 2. 遞迴
	var successor *ListNode
	return reverseBetweenHelper(head, left, right, &successor)
}

func reverseBetweenHelper(head *ListNode, m int, n int, successor **ListNode) *ListNode {
	// base case，找到要反轉的位置
	if m == 1 {
		return reverseN(head, n, successor)
	}
	// 如果 m != 1，head 就會一直前進，直到找到要反轉的範圍
	head.Next = reverseBetweenHelper(head.Next, m-1, n-1, successor)
	return head
}

func reverseN(head *ListNode, n int, successor **ListNode) *ListNode {
	if n == 1 {
		*successor = head.Next
		return head
	}
	last := reverseN(head.Next, n-1, successor)
	// left = 2, right = 4
	// 1 -> 2 -> 3 -> 4 -> 5
	// n = 1, head = 4, successor = 5
	// n = 2, last = 4, head = 3, 4 -> 3 -> 5
	// n = 3, last = 4, head = 2, 4 -> 3 -> 2 -> 5
	head.Next.Next = head
	head.Next = *successor
	return last
}
```

```java
import java.util.List;

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
    private ListNode successor;

    // Method 1. 遍歷
    // public ListNode reverseBetween(ListNode head, int left, int right) {
    //     ListNode dummy = new ListNode(0);
    //     dummy.next = head;
    //     ListNode p = dummy;
    //     for (int i = 0; i < left - 1; i++) {
    //         p = p.next;
    //     }

    //     ListNode prev = null;
    //     ListNode curr = p.next;
    //     for (int i = 0; i < right - left + 1; i++) {
    //         ListNode next = curr.next;
    //         curr.next = prev;
    //         prev = curr;
    //         curr = next;
    //     }
    //     p.next.next = curr;
    //     p.next = prev;
    //     return dummy.next;
    // }

    // Method 2. 遞迴
    public ListNode reverseBetween(ListNode head, int left, int right) {
        if (left == 1) {
            return reverseN(head, right);
        }
        head.next = reverseBetween(head.next, left - 1, right - 1);
        return head;
    }

    private ListNode reverseN(ListNode head, int n) {
        if (n == 1) {
            successor = head.next;
            return head;
        }
        ListNode last = reverseN(head.next, n - 1);
        head.next.next = head;
        head.next = successor;
        return last;
    }
}
```