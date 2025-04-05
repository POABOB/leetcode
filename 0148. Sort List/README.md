# Intuition

這題要將一個 linked list 進行排序，要求 `時間複雜度為 O(n log n)`，`空間複雜度為 O(1)`。

可以使用 `Merge Sort` 進行，不過有分成 `Top-Down` 與 `Bottom-Up`。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

## Method 1. Top-Down Merge Sort (遞迴)

- 定義一個找出 linked list `中間節點` 的函數，`Two Pointer 實作`。
- 將 linked list `遞迴切分`，切分到 `最小單位 (1 個元素)` 時，就可以將其進行 `遞迴合併`。
- 合併則是根據值的大小進行 `順序排列`。

```
Step 1. 4 -> 2 -> 1 -> 3
Step 2. 4 -> 2 , 1 -> 3   // 遞迴分割
Step 3. 4, 2, 1, 3        // 最小單位
Step 4. 2 -> 4 , 1 -> 3   // 遞迴順序合併
Step 5. 1 -> 2 -> 3 -> 4
```

## Method 2. Bottom-Up Merge Sort (迴圈)

- 先統計整個 linked list `長度 count`。
- 使用分組合併的方式，由小慢慢擴大 1, 2, 4, 8... 直到 `小於等於 count`
- 每一輪透過 `nextN` 函數切割，主要用來 return `下 n 個` linked list 的節點。
- merge 函數要記得 return tail，tail 指針用來記錄組與組之間的連結

```
Step 1. 4 -> 2 -> 5 -> 3 -> 6 -> 1 
Step 2. {2 -> 4}, 5 -> 3 -> 6 -> 1
Step 3. {2 -> 4} -> {3 -> 5}, 6 -> 1
Step 4. {2 -> 4} -> {3 -> 5} -> {1 -> 6}    // 1 個節點為一組排序完成
Step 5. {2 -> 3 -> 4 -> 5}, 1 -> 6
Step 6. {2 -> 3 -> 4 -> 5} -> {1 -> 6}
Step 7. 1 -> 2 -> 3 -> 4 -> 5 -> 6          // 2 個節點為一組排序完成
```

<!-- Describe your approach to solving the problem. -->

# Complexity

- Time complexity
    - Method 1. O(nlogn)，Merge Sort 固定時間複雜度。
    - Method 2. O(nlogn)，Merge Sort 固定時間複雜度。
- Space complexity
    - Method 1. O(logn) (call stack)
    - Method 2. O(1)

<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code

```go
package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// Method 1. Top Down
// func sortList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}

// 	mid := getMid(head)
// 	head2 := mid.Next
// 	mid.Next = nil
// 	return merge(sortList(head), sortList(head2))
// }

// func getMid(root *ListNode) *ListNode {
// 	fast, slow := root, root
// 	for fast.Next != nil && fast.Next.Next != nil {
// 		fast = fast.Next.Next
// 		slow = slow.Next
// 	}
// 	return slow
// }

// func merge(l1, l2 *ListNode) *ListNode {
// 	dummy := &ListNode{}
// 	curr := dummy
// 	for l1 != nil && l2 != nil {
// 		if l1.Val > l2.Val {
// 			curr.Next = l2
// 			l2 = l2.Next
// 		} else {
// 			curr.Next = l1
// 			l1 = l1.Next
// 		}
// 		curr = curr.Next
// 	}

// 	if l1 != nil {
// 		curr.Next = l1
// 	} else if l2 != nil {
// 		curr.Next = l2
// 	}
// 	return dummy.Next
// }

// Method 2. Bottom Up
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	listLen := 0
	curr := head
	for curr != nil {
		listLen++
		curr = curr.Next
	}

	dummy := &ListNode{Next: head}
	for i := 1; i <= listLen; i <<= 1 {
		curr = dummy.Next
		tail := dummy
		for curr != nil {
			left := curr
			right := nextN(left, i)
			curr = nextN(right, i)
			h, t := merge(left, right)
			// 因為 h 已經是斷開後的 sorted list 必須把被個別組別連接
			// E.g. {a1 <-> a2} {a3 -> a4 -> ... -> an-1 -> an}		  	=> 一開始各組為獨立的
			//    tail.Next tail
			//      {a1 -> a2} -> {a3 -> a4} {an-1 -> ... -> an}		=> tail.Next = h，產生新的組別後，讓組別連接
			//            tail tail.Next
			//      {a1 -> a2} -> {a3 -> a4} {an-1 -> ... -> an}		=> tail = t，接續往下傳遞
			//                 tail.Next tail
			tail.Next = h
			tail = t
		}

	}
	return dummy.Next
}

func nextN(root *ListNode, n int) *ListNode {
	if root == nil {
		return nil
	}
	for n > 1 && root.Next != nil {
		root = root.Next
		n--
	}
	next := root.Next
	root.Next = nil
	return next
}

func merge(l1, l2 *ListNode) (*ListNode, *ListNode) {
	dummy := &ListNode{}
	curr := dummy
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			curr.Next = l2
			l2 = l2.Next
		} else {
			curr.Next = l1
			l1 = l1.Next
		}
		curr = curr.Next
	}

	if l1 != nil {
		curr.Next = l1
	} else if l2 != nil {
		curr.Next = l2
	}

	// tail
	for curr.Next != nil {
		curr = curr.Next
	}
	return dummy.Next, curr
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

class Pair<T1, T2> {
    final T1 a;
    final T2 b;

    Pair(T1 a, T2 b) {
        this.a = a;
        this.b = b;
    }
}

class Solution {
    public ListNode sortList(ListNode head) {
        if (head == null) return null;

        // 計算整體 list 長度
        ListNode curr = head;
        int count = 0;
        while (curr != null) {
            curr = curr.next;
            count++;
        }

        // 每次進行分組，分組次數為 O(logn)，n 為 list 長度
        ListNode dummy = new ListNode();
        dummy.next = head;
        for (int i = 1; i <= count; i *= 2) {
            // 將組內 list 進行排序，O(n)
            curr = dummy.next;
            ListNode tail = dummy;
            while (curr != null) {
                ListNode left = curr;
                ListNode right = nextN(i, left);
                curr = nextN(i, right); // 剩下的 list
                Pair<ListNode, ListNode> sortedPartialList = merge(left, right);
                // tail 主要是用來連接 "組與組" 之間的
                tail.next = sortedPartialList.a;
                tail = sortedPartialList.b;
            }
        }
        return dummy.next;
    }

    private ListNode nextN(int n, ListNode root) {
        if (root == null) return null;
        while (n > 1 && root.next != null) {
            root = root.next;
            n--;
        }
        ListNode next = root.next;
        root.next = null;
        return next;
    }

    // sort 並合併兩 list，並 return head 與 tail
    private Pair<ListNode, ListNode> merge(ListNode l1, ListNode l2) {
        ListNode dummy = new ListNode();
        ListNode curr = dummy;
        while (l1 != null && l2 != null) {
            if (l1.val > l2.val) {
                curr.next = l2;
                l2 = l2.next;
            } else {
                curr.next = l1;
                l1 = l1.next;
            }
            curr = curr.next;
        }
        if (l1 != null) curr.next = l1;
        if (l2 != null) curr.next = l2;

        while (curr.next != null) {
            curr = curr.next;
        }
        return new Pair<>(dummy.next, curr);
    }
}
```