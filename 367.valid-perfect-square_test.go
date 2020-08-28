/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 *
 * https://leetcode-cn.com/problems/valid-perfect-square/description/
 *
 * algorithms
 * Easy (43.41%)
 * Likes:    164
 * Dislikes: 0
 * Total Accepted:    41.7K
 * Total Submissions: 96K
 * Testcase Example:  '16'
 *
 * 给定一个正整数 num，编写一个函数，如果 num 是一个完全平方数，则返回 True，否则返回 False。
 *
 * 说明：不要使用任何内置的库函数，如  sqrt。
 *
 * 示例 1：
 *
 * 输入：16
 * 输出：True
 *
 * 示例 2：
 *
 * 输入：14
 * 输出：False
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
func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}

	left, right := 0, num/2

	for left < right {
		mid := left + (right-left)/2
		if mid*mid == num {
			return true
		} else if mid*mid < num {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return false
}

// @lc code=end

func TestIsPerfectSquare(t *testing.T) {
	for index, tc := range []struct {
		num  int
		want bool
	}{
		{16, true},
		{14, false},
		{1, true},
	} {
		t.Run(fmt.Sprintf("case: %v", index), func(t *testing.T) {
			assert.Equal(t, tc.want, isPerfectSquare(tc.num))
		})
	}
}
