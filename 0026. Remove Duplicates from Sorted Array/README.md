# Intuition
這題是已經存在一個被排序過的 `陣列 nums`，需要把陣列中 `重複的元素刪除` 並 `將其他元素向前移動`，最後返回該 `全新不重複陣列的大小`。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 可以使用 `Two Pointer` 的方式來解決， slow pointer 為 `替換的指標`， fast pointer 為 `遍歷 nums 的指標`
- 遍歷時，要確定 `len(nums) > fast pointer`
  - 當 `nums[fast] != val` 時，就可以把 `nums[fast]` 替換到 `nums[slow]`
  - 每次替換判斷邏輯執行後， fast pointer 不斷的向前移動，直到 `nums[slow] != nums[fast]`
  - slow pointer 向前移動一個位置
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

func removeDuplicates(nums []int) int {
	fast, slow := 0, 0
	for len(nums) > fast {
		if nums[slow] != nums[fast] {
			nums[slow] = nums[fast]
		}

		for len(nums) > fast && nums[fast] == nums[slow] {
			fast++
		}
		slow++
	}
	return slow
}
```

```java
class Solution {
    public int removeDuplicates(int[] nums) {
        int fast = 0;
        int slow = 0;
        while (nums.length > fast) {
            if (nums[fast] != nums[slow]) {
                nums[slow] = nums[fast];
            }

            while (nums.length > fast && nums[fast] == nums[slow]) {
                fast++;
            }
            slow++;
        }
        return slow;
    }
}
```




