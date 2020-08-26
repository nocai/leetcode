/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 *
 * https://leetcode-cn.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (38.81%)
 * Likes:    484
 * Dislikes: 0
 * Total Accepted:    194.8K
 * Total Submissions: 501.8K
 * Testcase Example:  '4'
 *
 * 实现 int sqrt(int x) 函数。
 *
 * 计算并返回 x 的平方根，其中 x 是非负整数。
 *
 * 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
 *
 * 示例 1:
 *
 * 输入: 4
 * 输出: 2
 *
 *
 * 示例 2:
 *
 * 输入: 8
 * 输出: 2
 * 说明: 8 的平方根是 2.82842...,
 * 由于返回类型是整数，小数部分将被舍去。
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
func mySqrt(x int) int {
	left, right := 0, x
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			left = mid + 1
		} else if mid*mid > x {
			right = mid - 1
		}
	}
	return left - 1
}

// @lc code=end

func TestMySqrt(t *testing.T) {
	for index, tc := range []struct {
		in  int
		out int
	}{
		{4, 2},
		{8, 2},
		{9, 3},
		{10, 3},
	} {
		t.Run(fmt.Sprintf("case: %v", index), func(t *testing.T) {
			assert.Equal(t, tc.out, mySqrt(tc.in))
		})
	}
}
