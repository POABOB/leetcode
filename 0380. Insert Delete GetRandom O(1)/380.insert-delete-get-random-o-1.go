package leetcode

import "math/rand"

/*
 * @lc app=leetcode id=380 lang=golang
 *
 * [380] Insert Delete GetRandom O(1)
 */

// @lc code=start
type RandomizedSet struct {
	hashmap map[int]int
	arr     []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		hashmap: make(map[int]int),
		arr:     make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.hashmap[val]; ok {
		return false
	}
	this.hashmap[val] = len(this.arr)
	this.arr = append(this.arr, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if valIndex, ok := this.hashmap[val]; ok {
		lastIndex := len(this.arr) - 1

		// 把 hashmap 中 index 對調
		this.hashmap[this.arr[lastIndex]] = valIndex
		// 把 arr SWAP
		this.arr[valIndex], this.arr[lastIndex] = this.arr[lastIndex], this.arr[valIndex]
		this.arr = this.arr[:lastIndex]

		delete(this.hashmap, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end
