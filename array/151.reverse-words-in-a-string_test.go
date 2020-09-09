/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 翻转字符串里的单词
 *
 * https://leetcode-cn.com/problems/reverse-words-in-a-string/description/
 *
 * algorithms
 * Medium (43.09%)
 * Likes:    216
 * Dislikes: 0
 * Total Accepted:    93.7K
 * Total Submissions: 216.3K
 * Testcase Example:  '"the sky is blue"'
 *
 * 给定一个字符串，逐个翻转字符串中的每个单词。
 *
 *
 *
 * 示例 1：
 *
 * 输入: "the sky is blue"
 * 输出: "blue is sky the"
 *
 *
 * 示例 2：
 *
 * 输入: "  hello world!  "
 * 输出: "world! hello"
 * 解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
 *
 *
 * 示例 3：
 *
 * 输入: "a good   example"
 * 输出: "example good a"
 * 解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
 *
 *
 *
 *
 * 说明：
 *
 *
 * 无空格字符构成一个单词。
 * 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
 * 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
 *
 *
 *
 *
 * 进阶：
 *
 * 请选用 C 语言的用户尝试使用 O(1) 额外空间复杂度的原地解法。
 *
 */
package leetcode

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	ss := strings.Split(s, " ")
	sort.Sort(sort.Reverse(sort.StringSlice(ss)))
	return strings.Join(ss, " ")
}

// @lc code=end

func TestReverseWords(t *testing.T) {
	for index, tc := range []struct {
		input string
		want  string
	}{
		{" the sky is blue", "blue is sky the"},
		{"  hello world!  ", "world! hello"},
	} {
		t.Run(fmt.Sprintf("case: %v", index), func(t *testing.T) {
			assert.Equal(t, tc.want, reverseWords(tc.input))
		})
	}
}

func TestSortReverse(t *testing.T) {
	// ss := strings.Split("the sky is blue", " ")
	ss := []string{"the", "sky", "is", "blue"}
	sort.Sort(sort.StringSlice(ss))
	sort.Sort(sort.Reverse(sort.StringSlice(ss)))
	fmt.Println(ss)
}
