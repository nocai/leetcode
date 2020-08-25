package leetcode

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	for index, tc := range []TestCase{
		{123, 321},
		{-123, -321},
		{120, 21},
		{math.MaxInt32 + 1, 0},
		{math.MinInt32 - 1, 0},
	} {
		t.Run(tc.NameByIndex(index), func(t *testing.T) {
			assert.Equal(t, tc.out, reverse(tc.in.(int)))
		})
	}
}

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 *
 * https://leetcode-cn.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (34.62%)
 * Likes:    2110
 * Dislikes: 0
 * Total Accepted:    437.7K
 * Total Submissions: 1.3M
 * Testcase Example:  '123'
 *
 * 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
 *
 * 示例 1:
 *
 * 输入: 123
 * 输出: 321
 *
 *
 * 示例 2:
 *
 * 输入: -123
 * 输出: -321
 *
 *
 * 示例 3:
 *
 * 输入: 120
 * 输出: 21
 *
 *
 * 注意:
 *
 * 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回
 * 0。
 *
 */

// @lc code=start
func reverse(x int) int {
	var r int

	for ; x != 0; x /= 10 {
		r = r*10 + x%10
	}

	if r != int(int32(r)) {
		return 0
	}

	return r
}

// @lc code=end

// TestInt32 判断int32是否溢出
func TestInt32(t *testing.T) {
	i := math.MaxInt32
	assert.Equal(t, i, int(int32(i)))
	i += 1
	assert.NotEqual(t, i, int(int32(i)))

	y := math.MaxInt64
	assert.NotEqual(t, y, int(int32(y)))
}
