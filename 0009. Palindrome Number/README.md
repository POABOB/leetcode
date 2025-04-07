# Intuition

這題要我們判斷一個 `整數 x` 是否為 `「回文數 (Palindrome Number)」`，也就是 `正著讀` 和 `反著讀` 都一樣的數字。

例如：121、1221 是回文數；123、-121 則不是。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

- 如果 `x < 0`，直接回傳 `false`，因為負數有負號，不可能是回文數。
- 接著將數字 x 做反轉，並與原本的數字比較是否相等：
    - 每次取出 `x 的個位數 (x % 10)`，加到 `reverseX 的末端`。
    - 同時 x 要 `去掉個位數（x /= 10）`。
    - 最後比較反轉後的數字 `reverseX` 是否與 `原始數字` 相等即可。

<!-- Describe your approach to solving the problem. -->

# Complexity

- Time complexity
    - O(log₁₀x)，因為我們每次將 x 除以 10，最多會做 位數 次操作。
- Space complexity
    - O(1)，只用了常數額外變數。

<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code

```go
package leetcode

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	temp := x
	reverseX := 0
	for temp > 0 {
		reverseX = reverseX*10 + temp%10
		temp /= 10
	}
	return reverseX == x
}
```

```java
class Solution {
    public boolean isPalindrome(int x) {
        if (x < 0) return false;
        int temp = x;
        int reverseX = 0;
        while (temp > 0) {
            reverseX = reverseX * 10 + temp % 10;
            temp /= 10;
        }
        return reverseX == x;
    }
}
```