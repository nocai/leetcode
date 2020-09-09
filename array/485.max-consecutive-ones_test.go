/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] 最大连续1的个数
 *
 * https://leetcode-cn.com/problems/max-consecutive-ones/description/
 *
 * algorithms
 * Easy (56.67%)
 * Likes:    119
 * Dislikes: 0
 * Total Accepted:    51.9K
 * Total Submissions: 91.5K
 * Testcase Example:  '[1,0,1,1,0,1]'
 *
 * 给定一个二进制数组， 计算其中最大连续1的个数。
 *
 * 示例 1:
 *
 *
 * 输入: [1,1,0,1,1,1]
 * 输出: 3
 * 解释: 开头的两位和最后的三位都是连续1，所以最大连续1的个数是 3.
 *
 *
 * 注意：
 *
 *
 * 输入的数组只包含 0 和1。
 * 输入数组的长度是正整数，且不超过 10,000。
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
func findMaxConsecutiveOnes(nums []int) int {
	count, maxCount := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			count = 0
		}
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}

// @lc code=end

func TestFindMaxConsecutiveOnes(t *testing.T) {
	for index, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 0, 1, 1, 0, 1}, 2},
		{[]int{1, 1, 0, 1, 1, 1}, 3},
	} {
		t.Run(fmt.Sprintf("case: %v", index), func(t *testing.T) {
			assert.Equal(t, tc.want, findMaxConsecutiveOnes(tc.nums))
		})
	}
}
