/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] 判断子序列
 *
 * https://leetcode-cn.com/problems/is-subsequence/description/
 *
 * algorithms
 * Easy (50.51%)
 * Likes:    312
 * Dislikes: 0
 * Total Accepted:    85K
 * Total Submissions: 168.4K
 * Testcase Example:  '"abc"\n"ahbgdc"'
 *
 * 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
 *
 * 你可以认为 s 和 t 中仅包含英文小写字母。字符串 t 可能会很长（长度 ~= 500,000），而 s 是个短字符串（长度 <=100）。
 *
 *
 * 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
 *
 * 示例 1:
 * s = "abc", t = "ahbgdc"
 *
 * 返回 true.
 *
 * 示例 2:
 * s = "axc", t = "ahbgdc"
 *
 * 返回 false.
 *
 * 后续挑战 :
 *
 * 如果有大量输入的 S，称作S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T
 * 的子序列。在这种情况下，你会怎样改变代码？
 *
 * 致谢:
 *
 * 特别感谢 @pbrother 添加此问题并且创建所有测试用例。
 *
 */

package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func isSubsequence(s string, t string) bool {
	// double pointers

	sp, tp := 0, 0
	for sp < len(s) && tp < len(t) {
		if s[sp] == t[tp] {
			sp++
		}
		tp++
	}

	return sp == len(s)
}

// @lc code=end

func TestIsSubsequence(t *testing.T) {
	for index, tc := range []struct {
		s   string
		t   string
		rst bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
	} {
		t.Run(fmt.Sprintf("case: %v", index), func(t *testing.T) {
			assert.Equal(t, tc.rst, isSubsequence(tc.s, tc.t))
		})
	}
}
