/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 *
 * https://leetcode-cn.com/problems/daily-temperatures/description/
 *
 * algorithms
 * Medium (64.66%)
 * Total Accepted:    102.9K
 * Total Submissions: 159.2K
 * Testcase Example:  '[73,74,75,71,69,72,76,73]'
 *
 * 请根据每日 气温 列表，重新生成一个列表。对应位置的输出为：要想观测到更高的气温，至少需要等待的天数。如果气温在这之后都不会升高，请在该位置用 0
 * 来代替。
 *
 * 例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4,
 * 2, 1, 1, 0, 0]。
 *
 * 提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。
 *
 */
func dailyTemperatures(T []int) []int {
	rst := make([]int, len(T))
	stack := []int{}

	for i := 0; i < len(T); i++ {
		t := T[i]
		for len(stack) > 0 && t > T[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			rst[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	return rst
}
