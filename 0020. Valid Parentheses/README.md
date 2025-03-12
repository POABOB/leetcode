# Intuition
這題主要是透過 `堆疊 (Stack)` 的 LIFO(Last In First Out) 特性檢查括號是否成對匹配。我們可以使用 map 來存放對應的括號關係，並使用 stack 來幫助檢查是否符合條件。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 使用 `hashmap` 來存放括號的對應關係，E.g. `( 對應 )`，`[ 對應 ]`，`{ 對應 }`。
- 使用 stack 來儲存 `左括號`。
- 遍歷 s 字串：
	- 如果當前字元是 `左括號`，則將其 push。
	- 如果當前字元是 `右括號`：
		- 如果 stack `不為空`，且 stack `頂部的元素與當前右括號匹配`，則將 stack 頂部元素 pop。
		- 否則，return false。
- 最終檢查 stack `是否為空`，若為空則代表所有括號匹配成功。
<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity
    - O(n)，其中 n 是字串長度，每個字元最多被推入和彈出 stack 一次。
- Space complexity 
    - O(n)，最壞情況下 stack 可能會存儲所有的左括號。
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```go
package leetcode

func isValid(s string) bool {
	hashmap := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			if len(stack) > 0 && hashmap[stack[len(stack)-1]] == s[i] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
```
```java
class Solution {
	public boolean isValid(String s) {
		Stack<Character> stack = new Stack<>();

		for (int i = 0; i < s.length(); i++) {
			char c = s.charAt(i);
			if (c == '(' || c == '[' || c == '{') {
				stack.push(c);
			} else {
				if (stack.isEmpty()) {
					return false;
				}
				char top = stack.peek();
				if (c == ')' && top == '(' || c == ']' && top == '[' || c == '}' && top == '{') {
					stack.pop();
				} else {
					return false;
				}
			}
		}
		return stack.isEmpty();
	}
}
```