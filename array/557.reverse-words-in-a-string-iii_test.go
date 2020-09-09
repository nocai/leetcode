/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 *
 * https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/description/
 *
 * algorithms
 * Easy (71.47%)
 * Likes:    239
 * Dislikes: 0
 * Total Accepted:    94.6K
 * Total Submissions: 128.8K
 * Testcase Example:  `"Let's take LeetCode contest"`
 *
 * 给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
 *
 *
 *
 * 示例：
 *
 * 输入："Let's take LeetCode contest"
 * 输出："s'teL ekat edoCteeL tsetnoc"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。
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
func reverseWords_557(s string) string {
	return reverseWords_557_local(s)
}

// @lc code=end

func TestReverseWords_557(t *testing.T) {
	for index, tc := range []struct {
		input string
		want  string
	}{
		{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
	} {
		t.Run(fmt.Sprintf("case: %v", index), func(t *testing.T) {
			assert.Equal(t, tc.want, reverseWords_557_brute(tc.input))
		})
		t.Run(fmt.Sprintf("_557_local: %v", index), func(t *testing.T) {
			assert.Equal(t, tc.want, reverseWords_557_local(tc.input))
		})
	}
}

func reverseWords_557_local(s string) string {
	ss := []byte(s)

	for i := 0; i < len(ss); {
		idx := i
		for i < len(ss) && s[i] != ' ' {
			i++
		}

		left, right := idx, i-1
		for left < right {
			ss[left], ss[right] = ss[right], ss[left]
			left++
			right--
		}

		for i < len(s) && s[i] == ' ' {
			i++
		}
	}

	return string(ss)
}

func reverseWords_557_brute(s string) string {
	rst := []byte{}

	for i := 0; i < len(s); {
		idx := i
		for i < len(s) && s[i] != ' ' {
			i++
		}

		for j := i - 1; j >= idx; j-- {
			rst = append(rst, s[j])
		}

		for i < len(s) && s[i] == ' ' {
			rst = append(rst, ' ')
			i++
		}
	}

	return string(rst)
}
