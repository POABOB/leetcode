# Intuition
這題的目標是設計一個 `最小堆疊 (Min Stack)`，可以在常數時間內提供 `Push`、`Pop`、`Top` 和 `GetMin` 操作。我們需要使用 `兩個 stack` 來實現這些操作：一個儲存 `正常的元素`，另一個 `儲存當前最小值`。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 使用 stk 來儲存所有的元素。
- 使用 minStk 來儲存每個位置的最小值。每次 Push 一個元素時，我們會比較這個元素與當前最小值，將較小的值放入 minStk。
- 當執行 Pop 時，`同時從兩個 stack 中移除元素`。
- Top 會返回 stk 堆疊的頂端元素，GetMin 會返回 minStk 堆疊的頂端元素，即當前堆疊的最小值。

<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity
    - O(1)，每個操作 (Push、Pop、Top、GetMin) 都是常數時間。
- Space complexity 
    - O(n)，其中 n 為堆疊中的元素數量，stk 和 minStk 都儲存相同數量的元素。
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```go
package leetcode

type MinStack struct {
	stk    []int
	minStk []int
}

func Constructor() MinStack {
	return MinStack{stk: make([]int, 0), minStk: make([]int, 0)}
}

func (this *MinStack) Push(val int) {
	if len(this.minStk) == 0 || val <= this.minStk[len(this.minStk)-1] {
		this.minStk = append(this.minStk, val)
	} else {
		this.minStk = append(this.minStk, this.minStk[len(this.minStk)-1])
	}
	this.stk = append(this.stk, val)
}

func (this *MinStack) Pop() {
	if len(this.stk) > 0 {
		this.stk = this.stk[:len(this.stk)-1]
		this.minStk = this.minStk[:len(this.minStk)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stk[len(this.stk)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStk[len(this.minStk)-1]
}
```
```java
import java.util.Stack;

class MinStack {
    private final Stack<Integer> stk;
    private final Stack<Integer> minStk;

    public MinStack() {
        stk = new Stack<>();
        minStk = new Stack<>();
    }

    public void push(int val) {
        if (stk.isEmpty() || val <= minStk.peek()) {
            minStk.push(val);
        } else {
            minStk.push(minStk.peek());
        }
        stk.push(val);
    }

    public void pop() {
        stk.pop();
        minStk.pop();
    }

    public int top() {
        return stk.peek();
    }

    public int getMin() {
        return minStk.peek();
    }
}
```