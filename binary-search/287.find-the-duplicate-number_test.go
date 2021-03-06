/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 *
 * https://leetcode-cn.com/problems/find-the-duplicate-number/description/
 *
 * algorithms
 * Medium (65.84%)
 * Likes:    856
 * Dislikes: 0
 * Total Accepted:    95.1K
 * Total Submissions: 144.3K
 * Testcase Example:  '[1,3,4,2,2]'
 *
 * 给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和
 * n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。
 *
 * 示例 1:
 *
 * 输入: [1,3,4,2,2]
 * 输出: 2
 *
 *
 * 示例 2:
 *
 * 输入: [3,1,3,4,2]
 * 输出: 3
 *
 *
 * 说明：
 *
 *
 * 不能更改原数组（假设数组是只读的）。
 * 只能使用额外的 O(1) 的空间。
 * 时间复杂度小于 O(n^2) 。
 * 数组中只有一个重复的数字，但它可能不止重复出现一次。
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
func findDuplicate(nums []int) int {
	m := make(map[int]struct{})
	for _, num := range nums {
		if _, exist := m[num]; exist {
			return num
		} else {
			m[num] = struct{}{}
		}
	}
	return -1
}

// @lc code=end

func TestFindDuplicate(t *testing.T) {
	for index, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{3, 1, 3, 4, 2}, 3},
		{[]int{1, 3, 4, 2, 2}, 2},
	} {
		t.Run(fmt.Sprintf("case: %v", index), func(t *testing.T) {
			assert.Equal(t, tc.want, findDuplicate(tc.nums))
		})
	}
}
