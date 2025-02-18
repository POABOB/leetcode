# Intuition
這題給定一個陣列 height 裡面元素為直線的 `高度`，而題目希望可以找如果下滿雨在這 `兩個直線` 所 `擁有的面積中求最大值`。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 直線 `彼此的高度只能求最小值`，因爲另外一邊再怎麼高，`水還是會從低的地方漏出去`
- 寬度就是 Two Pointer `left` 與 `right` 之間的 `距離`
- 那如何決定要 `left 前進` 還是 `right 退後` 呢？
    - 只要 `比較高度` 就可以了
    - `height[right] > height[left]`， `left++`
    - `height[right] < height[left]`， `right--`
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

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0

	for right > left {
		hight := min(height[left], height[right])
		width := right - left
		area := hight * width
		if area > maxArea {
			maxArea = area
		}

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return maxArea
}
```
```java
class Solution {
    public int maxArea(int[] height) {
        int left = 0;
        int right = height.length - 1;
        int maxArea = 0;
        while (right > left) {
            int width = right - left;
            int hight = Math.min(height[right], height[left]);
            int area = width * hight;
            if (area > maxArea) {
                maxArea = area;
            }
            if (height[right] > height[left]) {
                left++;
            } else {
                right--;
            }
        }
        return maxArea;
    }
}
```