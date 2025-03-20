package leetcode

import "container/list"

/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */

// @lc code=start
type LRUCache struct {
	cap     int
	list    *list.List
	hashmap map[int]*list.Element
}

type KV struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		list:    list.New(),
		hashmap: make(map[int]*list.Element),
		cap:     capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	value, existed := this.hashmap[key]
	if !existed {
		return -1
	}

	this.moveExistedKeyToEnd(key)
	return value.Value.(KV).value
}

func (this *LRUCache) Put(key int, value int) {
	element, existed := this.hashmap[key]
	if existed {
		element.Value = KV{key, value}
		this.moveExistedKeyToEnd(key)
		return
	}

	newElement := this.list.PushBack(KV{key, value})
	this.hashmap[key] = newElement

	if len(this.hashmap) > this.cap {
		e := this.list.Front()
		delete(this.hashmap, e.Value.(KV).key)
		this.list.Remove(e)
	}
}

func (this *LRUCache) moveExistedKeyToEnd(key int) {
	var e *list.Element = this.hashmap[key]
	this.list.MoveToBack(e)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
