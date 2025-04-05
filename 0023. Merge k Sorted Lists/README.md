# Intuition

題目要求將 `k 個已排序` 的 linked list `合併` 成一個 `排序好` 的 linked list。

可以使用 `Min Heap` 的方式，每次從所有 linked list 的 `當前節點中選擇最小值`，並 `加入合併後的 list 中`。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

- 使用 Min Heap：
    - 建立一個 `MinHeap` 存放所有 linked list 的 `頭節點`。
    - 每次從 MinHeap 中 `取出最小節點`，並將 `該節點的下一個節點加入 heap 中`。
- 處理流程：
    - 初始化一個虛擬頭節點 dummy。
    - 當 heap 不為空時，`持續取出最小節點` 並 `連接到結果 list 中` 。
    - 若該節點 `有下一個節點` ，將其 `推入 heap 中`。
- 優點： 每次 heap 中最多 `維持 k 個節點`，能有效地從 k 條 lists 中 `找出最小值`，並 `依序構建合併`。

<!-- Describe your approach to solving the problem. -->

# Complexity

- Time complexity
    - O(nlogk)：總共會處理 n 個節點，每次 heap 操作為 log k，其中 k 為 linked list 數量。

<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity
    - O(k)：heap 最多存放 k 個節點。

<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code

```go
package leetcode

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	minHeap := make(MinHeap, 0)
	heap.Init(&minHeap)

	for _, list := range lists {
		if list != nil {
			heap.Push(&minHeap, list)
		}
	}

	dummy := &ListNode{}
	current := dummy
	for minHeap.Len() > 0 {
		element := heap.Pop(&minHeap).(*ListNode)
		current.Next = element
		// 由該 list 找最小值或找其他 list 中找出最小值
		if element.Next != nil {
			heap.Push(&minHeap, element.Next)
		}
		current = current.Next
	}
	return dummy.Next
}

type MinHeap []*ListNode

func (mh MinHeap) Less(i, j int) bool {
	return mh[i].Val < mh[j].Val
}

func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh MinHeap) Len() int {
	return len(mh)
}

func (mh *MinHeap) Push(x any) {
	*mh = append(*mh, x.(*ListNode))
}

func (mh *MinHeap) Pop() any {
	x := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return x
}
```

```java
import java.util.PriorityQueue;
import java.util.Queue;

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
    public ListNode mergeKLists(ListNode[] lists) {
        Queue<ListNode> queue = new PriorityQueue<>((o1, o2) -> o1.val - o2.val);
        for (ListNode node : lists) {
            if (node != null) queue.add(node);
        }

        ListNode dummy = new ListNode();
        ListNode curr = dummy;
        while (!queue.isEmpty()) {
            ListNode element = queue.poll();
            curr.next = element;
            if (element.next != null) queue.add(element.next);
            curr = curr.next;
        }
        return dummy.next;
    }
}
```