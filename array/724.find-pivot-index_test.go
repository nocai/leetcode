/*
 * @lc app=leetcode.cn id=724 lang=golang
 *
 * [724] 寻找数组的中心索引
 *
 * https://leetcode-cn.com/problems/find-pivot-index/description/
 *
 * algorithms
 * Easy (38.02%)
 * Likes:    215
 * Dislikes: 0
 * Total Accepted:    58K
 * Total Submissions: 151.1K
 * Testcase Example:  '[1,7,3,6,5,6]'
 *
 * 给定一个整数类型的数组 nums，请编写一个能够返回数组 “中心索引” 的方法。
 *
 * 我们是这样定义数组 中心索引 的：数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和。
 *
 * 如果数组不存在中心索引，那么我们应该返回 -1。如果数组有多个中心索引，那么我们应该返回最靠近左边的那一个。
 *
 *
 *
 * 示例 1：
 *
 * 输入：
 * nums = [1, 7, 3, 6, 5, 6]
 * 输出：3
 * 解释：
 * 索引 3 (nums[3] = 6) 的左侧数之和 (1 + 7 + 3 = 11)，与右侧数之和 (5 + 6 = 11) 相等。
 * 同时, 3 也是第一个符合要求的中心索引。
 *
 *
 * 示例 2：
 *
 * 输入：
 * nums = [1, 2, 3]
 * 输出：-1
 * 解释：
 * 数组中不存在满足此条件的中心索引。
 *
 *
 *
 * 说明：
 *
 *
 * nums 的长度范围为 [0, 10000]。
 * 任何一个 nums[i] 将会是一个范围在 [-1000, 1000]的整数。
 *
 *
 */
package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func pivotIndex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		sumLeft := sum(nums[:i])
		sumRight := sum(nums[i+1:])
		if sumLeft == sumRight {
			return i
		}
	}
	return -1
}

func sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func pivotIndex2(nums []int) int {
	sum, sumLeft := 0, 0
	for _, num := range nums {
		sum += num
	}

	for index, num := range nums {
		sumRight := sum - sumLeft - num
		if sumLeft == sumRight {
			return index
		}
		sumLeft += num
	}
	return -1
}

// @lc code=end

func TestPivotIndex(t *testing.T) {
	for index, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{-1, -1, -1, -1, -1, 0}, 2},
		{[]int{1, 7, 3, 6, 5, 6}, 3},
		{[]int{1, 2, 3}, -1},
	} {
		t.Run(fmt.Sprintf("case: %v", index), func(t *testing.T) {
			assert.Equal(t, tc.want, pivotIndex2(tc.nums))
		})
	}
}
