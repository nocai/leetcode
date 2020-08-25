package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	for index, tc := range []TestCase{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "rececar", "car"}, ""},
		{[]string{"a"}, "a"},
	} {
		t.Run(tc.NameByIndex(index), func(t *testing.T) {
			want := longestCommonPrefix(tc.in.([]string))
			assert.Equal(t, tc.out, want, tc.ErrorMessage(want))
		})
	}
}

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 *
 * https://leetcode-cn.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (38.63%)
 * Likes:    1223
 * Dislikes: 0
 * Total Accepted:    338.1K
 * Total Submissions: 875K
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 *
 * 如果不存在公共前缀，返回空字符串 ""。
 *
 * 示例 1:
 *
 * 输入: ["flower","flow","flight"]
 * 输出: "fl"
 *
 *
 * 示例 2:
 *
 * 输入: ["dog","racecar","car"]
 * 输出: ""
 * 解释: 输入不存在公共前缀。
 *
 *
 * 说明:
 *
 * 所有输入只包含小写字母 a-z 。
 *
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	strOne := strs[0]
	for i := 0; i < len(strOne); i++ {
		for j := 1; j < len(strs); j++ {
			strOther := strs[j]
			if i == len(strOther) || strOne[i] != strOther[i] {
				return strOne[:i]
			}
		}
	}
	return strOne
}

// @lc code=end

func TestRangeString(t *testing.T) {
	str := "abc"
	fmt.Printf("%T\n", str)
	for index, s := range str {
		fmt.Printf("%T, %T\n", s, str[index])
	}
}
