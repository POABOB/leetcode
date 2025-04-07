# Intuition

這題要求我們對一個 `非負整數陣列 digits` 進行 `加一` 操作，其中每個元素代表一個位數 `(高位在前)`，回傳 `加一後的新陣列`。
- 我們的目標是處理「進位」的情況，特別是連續 9 的情形，例如 [9, 9, 9] 會變成 [1, 0, 0, 0]。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

- 從陣列的 `最後一位` 開始向前遍歷：
  - 如果當前位數 `小於 9`，那就直接將它 `加一`，並且 `直接返回結果`。
  - 如果當前位數 `等於 9`，那就把它 `設為 0`，並繼續 `往前處理進位`。
- 如果全部位數都是 `9`，例如 `[9, 9, 9]`，那就會全部變成 `0, 0, 0`，這時候我們需要在 `最前面加上一個 1`，變成 `[1, 0, 0, 0]`。 

這個解法可以在原陣列上進行修改，只有在 `全是 9` 的情況下才會產生一個新陣列。

# Complexity

- Time complexity
    - O(n)，n 為陣列長度，需要從尾部最多遍歷到頭部。

<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity
    - O(1)，除了極少數情況下（全部為 9），不需要額外空間。

<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code

```go
package leetcode

func plusOne(digits []int) []int {
	lastIndex := len(digits) - 1

	for i := lastIndex; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	if digits[0] == 0 {
		digits[0] = 1
		digits = append(digits, 0)
	}
	return digits
}
```

```java
class Solution {
    public int[] plusOne(int[] digits) {
        int size = digits.length;
        for (int i = size - 1; i >= 0; i--) {
            if (digits[i] < 9) {
                digits[i]++;
                return digits;
            }
            digits[i] = 0;
        }
        if (digits[0] == 0) {
            digits[0] = 1;
            int[] newDigits = new int[size + 1];
            System.arraycopy(digits, 0, newDigits, 0, size);
            return newDigits;
        }
        return digits;
    }
}
```