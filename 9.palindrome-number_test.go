package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	for index, tc := range []TestCase{
		{121, true},
		{-121, false},
		{10, false},
		{3, true},
	} {
		t.Run(tc.NameByIndex(index), func(t *testing.T) {
			assert.Equal(t, tc.out, isPalindrome(tc.in.(int)))
		})
	}
}

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 *
 * https://leetcode-cn.com/problems/palindrome-number/description/
 *
 * algorithms
 * Easy (58.48%)
 * Likes:    1188
 * Dislikes: 0
 * Total Accepted:    424.4K
 * Total Submissions: 725.7K
 * Testcase Example:  '121'
 *
 * 判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
 *
 * 示例 1:
 *
 * 输入: 121
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: -121
 * 输出: false
 * 解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
 *
 *
 * 示例 3:
 *
 * 输入: 10
 * 输出: false
 * 解释: 从右向左读, 为 01 。因此它不是一个回文数。
 *
 *
 * 进阶:
 *
 * 你能不将整数转为字符串来解决这个问题吗？
 *
 */

// @lc code=start
func isPalindrome(x int) bool {
	// 特殊情况：
	// 1. x < 0, false
	if x < 0 {
		return false
	}

	// 2. x == 0, true
	if x == 0 {
		return true
	}

	// 3. x > 0 && 最后一位为0(x % 10 == 0), false
	if x%10 == 0 {
		return false
	}

	// 反转一半数字
	reverted := 0
	for ; x > reverted; x /= 10 {
		reverted = reverted*10 + x%10
	}

	// 回文数判断：
	// 1. 偶数长度时，x == reverted
	// 2. 奇数长度时，x == reverted / 10
	return x == reverted || x == reverted/10
}

// @lc code=end
