# Intuition
這題就是讓兩個 `已經排序好的` 陣列進行 `合併`，合併後狀態也是要 `排序過的`。通常一般的作法就是兩個陣列 `加在一起再排序`，但是也可以使用 `two pointer` 的方式來進行合併。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 我們明確知道 `nums2` 要被合併到 `nums1`，所以可以根據 `m-1` 和 `n-1` 當作兩者的最後 `index`
- 那麼我們就可以對兩陣列的 `index` 進行比大小
	- 當 `m != 0 && nums1[m-1] > nums2[n-1]` 時，代表 `nums1 的最大值大於 nums2 的最大值`，所以 `nums1 的最大值` 要搬到新的`後面的位置`(即`nums1[m +n - 1]`)
	- 並且要把 m 進行遞減，避免重複判斷，反之
- 這樣重複操作直到 `n 歸零`，陣列才算合併完成。
<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity
    - O(n)
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity 
    - O(1)
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```go
package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 把 nums2 塞入 nums1
	for n != 0 {
		// 如果兩元素倒數 m、n 筆，nums1 比較大，代表 nums1[m - 1] 的元素要放到最後面 nums1[m + n - 1]
		// 並且 m--，避免重複判斷，反之
		if m != 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}
```




