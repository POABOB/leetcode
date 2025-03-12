# Intuition
這題要將數學算式進行運算，該表達式包含 `加法`、`減法` 以及 `括號`。計算過程要 `遵循標準的數學運算優先級規則`。

# Approach

- 可以使用 `Stack` 來解決這個問題。遍歷表達式時，根據遇到的符號來處理數字和運算符：
    - 當遇到 `數字` 時，將其 `累積為一個完整的數字`。
    - 當遇到 `運算符` 時，根據 `前一個運算符` 處理數字並 `push`。
    - 當遇到 `括號` 時，`遞迴處理括號內的表達式`，確保括號內的計算優先於外部計算。

# Complexity
- Time complexity:
    - O(n)，其中 n 是表達式的長度。每個字符最多被處理一次。
- Space complexity:
    - O(n)，最壞情況下，我們可能需要儲存所有數字和運算符。

# Code

```go
package leetcode

import (
	"fmt"
	"unicode"
)

func calculate(s string) int {
	sum, _ := fn(s)
	return sum
}

func fn(s string) (int, string) {
	stk := make([]int, 0)
	var sign byte = '+'
	num := 0
	for len(s) > 0 {
		element := s[0]
		s = s[1:]
		if element == '(' {
			num, s = fn(s)
		}
		if unicode.IsDigit(rune(element)) {
			num = num*10 + int(element-'0')
		}

		if (!unicode.IsDigit(rune(element)) && element != ' ') || len(s) == 0 {
			switch sign {
			case '+':
				stk = append(stk, num)
			case '-':
				stk = append(stk, -num)
			case '*':
				stk[len(stk)-1] *= num
			case '/':
				stk[len(stk)-1] /= num
			}
			sign = element
			num = 0
		}
		if element == ')' {
			break
		}
	}
	sum := 0
	for _, v := range stk {
		sum += v
	}
	return sum, s
}
```

```java
import java.util.Stack;

class Solution {
    private String proccessString;

    public int calculate(String s) {
        proccessString = s;
        return fn();
    }

    private int fn() {
        Stack<Integer> stk = new Stack<>();
        int num = 0;
        char op = '+';
        while (!proccessString.isEmpty()) {
            char element = proccessString.charAt(0);
            proccessString = proccessString.substring(1);

            if (element == '(') {
                num = fn();
            }
            if (Character.isDigit(element)) {
                num = num * 10 + Character.getNumericValue(element);
            }

            if ((!Character.isDigit(element) && element != ' ') || proccessString.isEmpty()) {
                int last;
                switch (op) {
                    case '+':
                        stk.push(num);
                        break;
                    case '-':
                        stk.push(-num);
                        break;
                    case '*':
                        last = stk.pop();
                        stk.push(last * num);
                        break;
                    case '/':
                        last = stk.pop();
                        stk.push(last / num);
                        break;
                }
                num = 0;
                op = element;
            }
            if (element == ')') {
                break;
            }
        }

        int sum = 0;
        while (!stk.isEmpty()) {
            sum += stk.pop();
        }
        return sum;
    }
}
```