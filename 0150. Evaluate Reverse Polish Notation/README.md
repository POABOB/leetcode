# Intuition
這題目要求我們評估一個 `逆波蘭表示法 (Reverse Polish Notation, RPN)` 的數學表達式，並返回結果。逆波蘭表示法的特點是 `操作符` 在運算元之後出現，這使得我們可以利用堆疊來簡單地計算結果。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 使用堆疊來 `儲存數字`和 `中間計算` 的結果。
- 當遇到 `數字` 時，將其 `push`。
- 當遇到 `操作符 (+, -, *, /)` 時，從 stack 中 `pop` 一個數字，並與 stack 內 `最後元素` 進行運算。
最後，堆疊中唯一剩下的數字即為最終結果。

<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity
    - O(n)，其中 n 為 tokens 陣列的長度，每個元素最多只被處理一次。
- Space complexity
	- O(n)，堆疊的大小最多為 tokens 陣列的長度。
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```go
package leetcode

import "strconv"

func evalRPN(tokens []string) int {
	stk := make([]int, 0, len(tokens))

	for _, token := range tokens {
		switch token {
		case "+":
			stk[len(stk)-2] += stk[len(stk)-1]
			stk = stk[:len(stk)-1]
		case "-":
			stk[len(stk)-2] -= stk[len(stk)-1]
			stk = stk[:len(stk)-1]
		case "/":
			stk[len(stk)-2] /= stk[len(stk)-1]
			stk = stk[:len(stk)-1]
		case "*":
			stk[len(stk)-2] *= stk[len(stk)-1]
			stk = stk[:len(stk)-1]
		default:
			num, _ := strconv.Atoi(token)
			stk = append(stk, num)
		}
	}
	return stk[0]
}
```
```java
import java.util.Stack;

class Solution {
    public int evalRPN(String[] tokens) {
        Stack<Integer> stk = new Stack<>();
        for (String token : tokens) {
            int temp;
            switch (token) {
                case "+":
                    temp = stk.pop();
                    stk.push(stk.pop() + temp);
                    break;
                case "-":
                    temp = stk.pop();
                    stk.push(stk.pop() - temp);
                    break;
                case "*":
                    temp = stk.pop();
                    stk.push(stk.pop() * temp);
                    break;
                case "/":
                    temp = stk.pop();
                    stk.push(stk.pop() / temp);
                    break;
                default:
                    stk.push(Integer.parseInt(token));
            }
        }
        return stk.pop();
    }
}
```