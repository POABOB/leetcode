package leetcode

/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

// @lc code=start
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

// @lc code=end
