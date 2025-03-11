# Intuition

這題要判斷一個數字是否為 Happy Number，即對數字的 `每位數字進行平方和運算`，最終是否能`收斂到 1`。若進入 `無窮迴圈` 則代表不是 Happy Number。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

- 可以使用快慢指針(Floyd's Cycle Detection Algorithm) 來解決這個問題：
    - 慢指針 (slow) 每次計算一次平方和。
    - 快指針 (fast) 每次計算兩次平方和。
    - `若兩者相遇`，則可能 `形成循環`，若最終 slow 指針等於 1，則代表是 Happy Number。
<!-- Describe your approach to solving the problem. -->

# Complexity

- Time complexity
    - O(logn)

<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity
    - O(1)

<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code

```go
package leetcode

func isHappy(n int) bool {
	slow, fast := n, n
	slow = getSum(slow)
	fast = getSum(getSum(fast))

	for slow != fast {
		slow = getSum(slow)
		fast = getSum(getSum(fast))
	}
	return slow == 1
}

func getSum(n int) int {
	var sum int
	for n > 0 {
		remainder := n % 10
		sum += remainder * remainder
		n /= 10
	}
	return sum
}
```

```java
class Solution {
    public boolean isHappy(int n) {
        int slow = n, fast = n;
        slow = getSum(slow);
        fast = getSum(getSum(fast));
        while (slow != fast) {
            slow = getSum(slow);
            fast = getSum(getSum(fast));
        }
        return slow == 1;
    }

    private int getSum(int n) {
        int sum = 0;
        while (n > 0) {
            int remainder = n % 10;
            sum += remainder * remainder;
            n /= 10;
        }
        return sum;
    }
}
```
