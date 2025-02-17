# Intuition
這題要把整數轉換成羅馬數字，通常就是把羅馬數字 `從大扣到小 (M -> I)`，其中要特別也是要特別注意，IV, IX, XL, ... 這類型的值。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 建立一個 MAP，根據阿拉伯數字對應相對應的羅馬數字
- 建立一個陣列，預設值為 [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1]
- 使用 foreach 來讓陣列由大到小不斷地 `將 num 進行扣除`
	- 如果 `convertMap[val] < num`，就代表數字不夠扣，要 `往下遞減`，`直到 num 歸零`
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

func intToRoman(num int) string {
	var sb strings.Builder
	convertMap := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	// 從大到小
	iterator := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	for _, val := range iterator {
		for num >= val {
			sb.WriteString(convertMap[val])
			num -= val
		}
	}
	return sb.String()
}
```
```java
class Solution {
    public String intToRoman(int num) {
        Map<Integer, String> currentMap = getConvertMap();
        List<Integer> iterator = new ArrayList<>(currentMap.keySet());
        iterator = iterator.stream().sorted(Collections.reverseOrder()).collect(Collectors.toList());
        StringBuilder sb = new StringBuilder();
        for (int currNum : iterator) {
            while (num >= currNum) {
                sb.append(currentMap.get(currNum));
                num -= currNum;
            }
        }
        return sb.toString();
    }

    private Map<Integer, String> getConvertMap() {
        Map<Integer, String> currentMap = new HashMap<>();
        currentMap.put(1, "I");
        currentMap.put(4, "IV");
        currentMap.put(5, "V");
        currentMap.put(9, "IX");
        currentMap.put(10, "X");
        currentMap.put(40, "XL");
        currentMap.put(50, "L");
        currentMap.put(90, "XC");
        currentMap.put(100, "C");
        currentMap.put(400, "CD");
        currentMap.put(500, "D");
        currentMap.put(900, "CM");
        currentMap.put(1000, "M");
        return currentMap;
    }
}
```