# Intuition
本題為判斷 `單向 Linked List` 是否有 `Cycle` 的問題。給定一個 `Linked List 的頭節點 head`，需要確認是否存在 `某個節點能夠再次被訪問到`。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 可以使用 快慢指針（Floyd's Cycle Detection Algorithm, 龍龜賽跑演算法） 來解決。
- 設定兩個指針 `slow` 和 `fast`，slow 每次 `移動一步`，fast 每次 `移動兩步`。
- 如果 fast 指針能夠追上 slow 指針（即 `fast == slow`），表示 Linked List `存在 Cycle`，return true。
- 如果 fast 指針到達 `nil`，表示 Linked List `無Cycle`，return false。

<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity
    - O(n)，每個節點最多被訪問兩次，一次是 slow，一次是 fast。
- Space complexity
    - O(1)，只使用了兩個指針變數，沒有額外的記憶體使用。
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```go
package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}
```

```java
class ListNode {
  int val;
  ListNode next;

  ListNode(int x) {
    val = x;
    next = null;
  }
}

public class Solution {
  public boolean hasCycle(ListNode head) {
    ListNode fast = head, slow = head;
    while (fast != null && fast.next != null) {
      fast = fast.next.next;
      slow = slow.next;
      if (fast == slow) {
        return true;
      }
    }
    return false;
  }
}
```