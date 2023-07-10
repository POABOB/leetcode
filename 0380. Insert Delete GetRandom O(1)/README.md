# Intuition
基本上就是實作一個 `不可以重複的 Set`，然後進行基本操作，其中要注意的是他在隨機存取的時候需要的 `時間複雜度為 O(1)`。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 定義 `hashmap` 和 `arr`，`hashmap` 主要負責儲存`值`和`陣列索引`，`arr` 就是位隨機存取而設計的
<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity
    - O(1)
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity 
    - O(n)
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```go
package leetcode

import "math/rand"

type RandomizedSet struct {
	hashmap map[int]int
	arr     []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		hashmap: make(map[int]int),
		arr:     []int{},
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
	if val_index, ok := this.hashmap[val]; ok {
		last_index := len(this.arr) - 1

		// 把 hashmap 中 index 對調
		this.hashmap[this.arr[last_index]] = val_index
		// 把 arr SWAP
		this.arr[val_index], this.arr[last_index] = this.arr[last_index], this.arr[val_index]
		this.arr = this.arr[:last_index]

		delete(this.hashmap, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]
}
```