# Intuition
將兩個已經被排序過的 `Linked-List` 進行合併，主要的難度就是如何控制指標來 `指向下一個節點`，可以使用 `遍歷` 或是 `遞迴` 來解決，通常使用 `遞迴` 會比較好懂也比較快。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- `Method 1.` (遍歷)
	1. 建立一個空的 `Linked-List` 叫做 `head`，用作答案返回
	2. 再建立一個 `dummy` 負責當作指標指向當前的 `head`
	3. 使用迴圈來 `比大小`，較小的將會被指向到 `dummy.Next`，陸續操作到 `其中一個` Linked-List 為空
	4. 判斷是不是還有`剩下的 Linked-List 沒有被指定`，並將 `dummy.Next` 指向它
- `Method 2.` (遞迴)
	1. 設定終止條件，當 `list1 == nil` 返回 `list2`，反之
	2. `比大小`，如果誰比較小，那就將小的 `Linked-List` 的 `Next` 指向 `mergeTwoLists()`，並返回小的 `Linked-List`
		- e.g. 假設 `list2` 小，那就會是 `list2.Next = mergeTwoLists(list1, list2.Next)`，反之
- 參考資料 
    - https://labuladong.github.io/algo/di-ling-zh-bfe1b/shuang-zhi-0f7cc/
<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity
    - Method 1. O(n)
    - Method 2. O(n)(call stack)
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity 
    - Method 1. O(1)
    - Method 2. O(1)(call stack)
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```go
package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Method 1. 遍歷
	// 建立一個 head 作為返回的答案
	// head := &ListNode{-1, nil}
	// dummy 是一個暫存指標
	// dummy := head

	// for list1 != nil && list2 != nil {
	// 	if list1.Val > list2.Val {
	// 		dummy.Next = list2
	// 		list2 = list2.Next
	// 	} else {
	// 		dummy.Next = list1
	// 		list1 = list1.Next
	// 	}
	// 	dummy = dummy.Next
	// }

	// if list1 != nil {
	// 	dummy.Next = list1
	// }

	// if list2 != nil {
	// 	dummy.Next = list2
	// }
	// return head.Next

	// Method 2. 遞迴
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val > list2.Val {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}

	list1.Next = mergeTwoLists(list1.Next, list2)
	return list1
}
```