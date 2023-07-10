# Intuition
TODO

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
TODO
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

import "sort"

func hIndex(citations []int) int {
	c_len := len(citations)
	sort.Ints(citations)

	for _, v := range citations {
		if c_len > v {
			c_len--
		}
	}

	return c_len
}
```