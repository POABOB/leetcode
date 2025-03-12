# Intuition
這題主要是要 `簡化 Unix 檔案路徑`，透過 `字串處理` 與 `堆疊 (Stack)` 來達成目標。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 先將 path 以 `/` 作為分隔符號，`分割成 parts 陣列`。
- 使用 stack 來模擬路徑處理：
	- 遇到 `.` 或 `空字串`，直接`忽略`。
	- 遇到 `..`，代表要 `返回上一層`，如果 stack 不為空，則 `pop stack 最上層的元素`。
	- 遇到 `其他字串`，則將其視為有效目錄，push stack。
	- 最後檢查 stack，`空則回傳 /`，不然則將 stack 組合。

# Complexity
- Time complexity
    - O(n)，其中 n 為 path 的長度，每個部分最多只會被處理一次。
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity 
    - O(n)，最壞情況下 stack 會存放所有有效的路徑部分。
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```go
package leetcode

import "strings"

func simplifyPath(path string) string {
	stack := make([]string, 0, len(path))
	parts := strings.Split(path, "/")
	for _, part := range parts {
		if part == "" || part == "." {
			continue
		} else if part == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, part)
		}
	}
	if len(stack) == 0 {
		return "/"
	}
	return "/" + strings.Join(stack, "/")
}
```

```java
import java.util.Stack;

class Solution {
	public String simplifyPath(String path) {
		String[] parts = path.split("/");
		Stack<String> stack = new Stack<>();
		for (String part : parts) {
			if (part.equals(".") || part.isEmpty()) {
				continue;
			} else if (part.equals("..")) {
				if (!stack.isEmpty()) {
					stack.pop();
				}
			} else {
				stack.push(part);
			}
		}
		StringBuilder sb = new StringBuilder();
		while (!stack.isEmpty()) {
			sb.insert(0, "/" + stack.pop());
		}
		return sb.isEmpty() ? "/" : sb.toString();
	}
}
```