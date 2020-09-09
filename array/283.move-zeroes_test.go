/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 *
 * https://leetcode-cn.com/problems/move-zeroes/description/
 *
 * algorithms
 * Easy (62.05%)
 * Likes:    706
 * Dislikes: 0
 * Total Accepted:    198.5K
 * Total Submissions: 319.4K
 * Testcase Example:  '[0,1,0,3,12]'
 *
 * 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
 *
 * 示例:
 *
 * 输入: [0,1,0,3,12]
 * 输出: [1,3,12,0,0]
 *
 * 说明:
 *
 *
 * 必须在原数组上操作，不能拷贝额外的数组。
 * 尽量减少操作次数。
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
func moveZeroes(nums []int) {
	slow, fast := 0, 0

	for ; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}

// @lc code=end

func TestMoveZeroes(t *testing.T) {
	for index, tc := range []struct {
		nums []int
		want []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{1, 1, 0, 3, 12}, []int{1, 1, 3, 12, 0}},
	} {
		t.Run(fmt.Sprintf("case: %v", index), func(t *testing.T) {
			moveZeroes(tc.nums)
			assert.Equal(t, tc.want, tc.nums)
		})
	}
}
