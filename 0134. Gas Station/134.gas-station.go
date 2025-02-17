package leetcode

/*
 * @lc app=leetcode id=134 lang=golang
 *
 * [134] Gas Station
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	sum := 0
	minSum := 0
	start := 0
	for i := 0; i < n; i++ {
		sum += gas[i] - cost[i]
		// 行駛途中發現油不夠前往下一站，要換地點上車
		if minSum > sum {
			// 在最耗油的一站開始，才會是最佳解
			start = i + 1
			minSum = sum
		}
	}

	// 到最後油量都是負的
	if sum < 0 {
		return -1
	}

	if start == n {
		return 0
	}
	return start
}

// @lc code=end
