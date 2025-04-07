# Intuition
這題要我們找出一個`整數 n` 的 `階乘(n!)` 中，尾端 `有多少個連續的 0`。

由於 0 是由 `2 * 5` 得來的，而在階乘中，2 的數量 `遠遠多於` 5，所以我們只需要 `數有多少個 5` 作為因數即可。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

- 我們要統計在 1 到 n 這些數字中，`有多少個數字是 5 的倍數`：
  - 每個 5 的倍數 `貢獻一個 5`。
  - 每個 25 的倍數 `會額外再貢獻一個`（因為 25 = `5 * 5`）。
  - 每個 125 的倍數 `又會額外多一個`，`以此類推`。
- 因此，我們可以不斷地用 `n / 5`、`n / 25`、`n / 125`... 相加，直到 `除數大於 n` 為止。

這種做法比暴力計算 n! 快非常多，因為不需要真正計算階乘，只需要做除法操作。

<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity
    - O(log₅n)，因為每次除以 5，最多會除 log₅n 次。
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity 
    - O(1)，只用了常數空間。
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```go
package leetcode

func trailingZeroes(n int) int {
	// n < 10^4
	// 5   的倍數: 5, 10, 15, 20, 25, 30, 35    -> 提供1個 5
	// 25  的倍數: 25, 50, 75, ...              -> 提供2個 5
	// 125 的倍數: 125, 250, 375, ...           -> 提供3個 5

	// E.g. n = 10
	// 10 / 5 = 2，會有 2 個 5
	// 總共 2 個
	// E.g. n = 30
	// 30 / 5  = 6，會有 6 個 5
	// 30 / 25 = 1，會有 1 個 5 (25 會提供 2 個，所以要補回來)
	// 總共 7 個
	// E.g. n = 130
	// 130 / 5 = 26，會有 26 個 5
	// 130 / 25 = 5，會有 5 個 5 (25 會提供 2 個，所以要補回來)
	// 130 / 125 = 1，會有 1 個 5 (125 會提供 3 個，所以要補回來)
	// 總共 32 個
	res := 0
	divisor := 5
	for n >= divisor {
		res += n / divisor
		divisor *= 5
	}
	return res
}
```

```java
class Solution {
    public int trailingZeroes(int n) {
        int res = 0;
        int divider = 5;
        while (n >= divider) {
            res += n / divider;
            divider *= 5;
        }
        return res;
    }
}
```
