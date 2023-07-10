# Intuition
這題主要是給定一個陣列，透過陣列中 n 筆元素，找到出現次數`大於 n / 2` 的元素，最基本的做法可以使用 `hashmap` 去紀錄最多筆的元素，也可以使用排序先整理陣列，然後再去取中間值 `nums[len(nums)/2]`。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- `Method 1.` 透過 HashMap 儲存 `nums 個別元素的出現次數` (`map[元素]count`)・再透過迴圈尋找符合條件的 `眾數`
- `Method 2.` (公式) 使用 `Boyer–Moore majority vote algorithm(摩爾投票算法)` 來解題
    - 主要原理透過 `相鄰的元素` 進行比較，若不相等，將它們從陣列中 `去除`，如此一來最後 `剩下的`，就會是 `最多數的`，因為有 `超過半數的保證`
- `Method 3.` `排序`，因為一個陣列出現次數要 `大於 n / 2`，代表排序後它一定會是 `中位數`，所以再 `取中間值` 就是答案
- 參考資料
    - https://www.geeksforgeeks.org/majority-element/
    - https://ithelp.ithome.com.tw/articles/10213285
<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity
    - Method 1. O(n)
    - Method 2. O(n)
    - Method 3. O(nlogn)
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity 
    - Method 1. O(n)
    - Method 2. O(1)
    - Method 3. O(1)
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```go
package leetcode

// import "sort"

func majorityElement(nums []int) int {

	// Method 1. HashMap 計數，最基本的寫法
	// nums_len, ans := len(nums), 0
	// count := make(map[int]int, nums_len)
	// for _, v := range nums {
	// 	count[v]++
	// }

	// for _, v := range nums {
	// 	if count[v] > nums_len/2 && count[v] > ans {
	// 		ans = v
	// 	}
	// }
	// return ans

	// Method 2. Boyer–Moore majority vote algorithm(摩爾投票算法)
	target, count, nums_len := 0, 0, len(nums)
	for i := 0; i < nums_len; i++ {
		if count == 0 {
			// 當 count== 0，假設 nums[i] 是眾數，並出現一次
			target = nums[i]
			count = 1
		} else if nums[i] == target {
			// 遇到一樣就累加
			count++
		} else {
			// 沒有就遞減
			count--
		}
	}
	// 通過 ++/--，此時的count會大於0，返回的 target 一定會是目標眾數
	return target

	// Method 3. 先排序，取最中間的元素就是答案，因為出現次數最多，必然在排序後會在中間
	// sort.Ints(nums)
	// return nums[len(nums)/2]
}
```




