# Intuition

本題要求實作 `LRU (Least Recently Used) Cache`，即 `最近最少使用淘汰機制的快取`。

- 實作一個具有 `固定容量 capacity` 的 LRUCache，需要支援以下兩個操作：
  - `Get(key int) int`：取得`快取中的資料`，若 key 存在，則返回對應的 value，否則返回 -1。
  - `Put(key int, value int)`：`存入資料`，若 key 已存在，則更新 value，若 key 不存在，則插入新的 key-value，但當 `快取容量已滿` 時，`需移除最久未使用的資料` 來騰出空間。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

- 為了有效率地實作 LRU Cache，我們需要：
  - Hashmap：O(1) 時間內 `快速查找 key` 是否存在。
  - Doubly Linked List：O(1) 時間內實現節點的 `插入`、`刪除`、`移動`。
- Doubly Linked List 的作用是：
  - `快速訪問最近使用的元素`（將新存入或被存取的 key 移到尾部）。
  - `刪除最久未使用的元素`（從頭部刪除）。
## 具體步驟
1. 使用 map[int]*list.Element 儲存 `key` 與 `Linked List 節點` 的對應關係，確保 Get 和 Put 操作的查找是 O(1) 時間複雜度。
2. 使用 container/list 的 list.List 來維護 LRU 的順序：
  - `PushBack` 來將最近存取的 key-value 移到尾部。
  - `MoveToBack` 來更新已有的 key-value。
  - `Front()` 來取得最久未使用的 key，以便刪除。
3. 實作 Get 方法：
  - 若 key 存在，將該節點移到 Linked List `尾部`，並返回 value。
  - 若 key 不存在，則返回 -1。
4. 實作 Put 方法：
  - 若 key 已存在，則更新值並移到 Linked List `尾部`。
  - 若 key 不存在，則新增節點至 Linked List `尾部`。
  - 若容量超過 capacity，則 刪除 Linked List `頭部`（最久未使用的節點），並從 hashmap 中 `移除對應的 key`。


<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity
  - Get O(1)：透過 map 快速查找，再透過 MoveToBack 來更新順序。
  - Put O(1)：透過 map 查找，並透過 PushBack 來插入新節點，若超過容量，則透過 Front() 來刪除最久未使用的節點。
- Space complexity
    - O(n)，因為 hashmap 和 list 需要儲存最多 capacity 個元素。
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```go
package leetcode


import "container/list"

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
```

```java
import java.util.LinkedHashMap;
import java.util.Map;

class LRUCache {

    private final int capacity;
    private final Map<Integer, Integer> cache = new LinkedHashMap<>();

    public LRUCache(int capacity) {
        this.capacity = capacity;
    }

    public int get(int key) {
        if (!cache.containsKey(key)) {
            return -1;
        }
        makeRecently(key);
        return cache.get(key);
    }

    public void put(int key, int val) {
        if (cache.containsKey(key)) {
            cache.put(key, val);
            makeRecently(key);
            return;
        }

        if (cache.size() >= this.capacity) {
            int oldestKey = cache.keySet().iterator().next();
            cache.remove(oldestKey);
        }
        cache.put(key, val);
    }

    private void makeRecently(int key) {
        int val = cache.get(key);
        cache.remove(key);
        cache.put(key, val);
    }
}
```