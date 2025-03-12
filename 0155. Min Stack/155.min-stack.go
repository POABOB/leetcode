package leetcode

/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */

// @lc code=start
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

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
