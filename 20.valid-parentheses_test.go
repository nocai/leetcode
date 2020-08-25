/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode-cn.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (42.84%)
 * Likes:    1807
 * Dislikes: 0
 * Total Accepted:    385.8K
 * Total Submissions: 900.4K
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 *
 * 有效字符串需满足：
 *
 *
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 *
 *
 * 注意空字符串可被认为是有效字符串。
 *
 * 示例 1:
 *
 * 输入: "()"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: "()[]{}"
 * 输出: true
 *
 *
 * 示例 3:
 *
 * 输入: "(]"
 * 输出: false
 *
 *
 * 示例 4:
 *
 * 输入: "([)]"
 * 输出: false
 *
 *
 * 示例 5:
 *
 * 输入: "{[]}"
 * 输出: true
 *
 */

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	for index, tc := range []TestCase{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"(", false},
		{")", false},
		{"()))", false},
	} {
		t.Run(tc.NameByIndex(index), func(t *testing.T) {
			got := isValid(tc.in.(string))
			assert.Equal(t, tc.out, got, tc.ErrorMessage(got))
		})
	}
}

// @lc code=start
func isValid(s string) bool {
	m := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	// 栈：先进后出
	stack := []rune{}

	for _, c := range s {
		if v, exist := m[c]; exist {
			// 前括号，入栈
			stack = append(stack, v)
		} else {
			// 后括号
			if len(stack) == 0 {
				// 没有匹配的前括括号，返回false
				return false
			}
			// 取最近一个前括号，进行匹配
			v = stack[len(stack)-1]
			if v != c {
				// 不匹配，返回false
				return false
			}
			// 匹配，出栈。进行下一轮循环
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// @lc code=end
