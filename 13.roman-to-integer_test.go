package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	for index, tc := range []TestCase{
		{"I", 1},
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"MDCXCV", 1695},
	} {
		t.Run(tc.NameByIndex(index), func(t *testing.T) {
			assert.Equal(t, tc.out, romanToInt(tc.in.(string)))
		})
	}
}

/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 *
 * https://leetcode-cn.com/problems/roman-to-integer/description/
 *
 * algorithms
 * Easy (62.01%)
 * Likes:    994
 * Dislikes: 0
 * Total Accepted:    241.5K
 * Total Submissions: 389.4K
 * Testcase Example:  '"III"'
 *
 * 罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
 *
 * 字符          数值
 * I             1
 * V             5
 * X             10
 * L             50
 * C             100
 * D             500
 * M             1000
 *
 * 例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V +
 * II 。
 *
 * 通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数
 * 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
 *
 *
 * I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
 * X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
 * C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
 *
 *
 * 给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。
 *
 * 示例 1:
 *
 * 输入: "III"
 * 输出: 3
 *
 * 示例 2:
 *
 * 输入: "IV"
 * 输出: 4
 *
 * 示例 3:
 *
 * 输入: "IX"
 * 输出: 9
 *
 * 示例 4:
 *
 * 输入: "LVIII"
 * 输出: 58
 * 解释: L = 50, V= 5, III = 3.
 *
 *
 * 示例 5:
 *
 * 输入: "MCMXCIV"
 * 输出: 1994
 * 解释: M = 1000, CM = 900, XC = 90, IV = 4.
 *
 */

// @lc code=start

func romanToInt(s string) int {
	// 思路
	// 从右到左遍历罗马数字的字符，将罗马字符映射为对应的阿拉伯数字，若当前的数字大于或等于前一个数字，则加，否则减。
	// 例如：XXVII等于1+1+5+10+10 = 27 、IX等于10-1=9、XCI等于1+100-10
	pre, result := 0, 0

	for i := len(s) - 1; i >= 0; i-- {
		cur := m[s[i]]
		if cur < pre {
			result -= cur
		} else {
			result += cur
		}
		pre = cur
	}

	return result

	// 通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例
	// 特例就是左边比右边小时，只加两数之差
	// for i := 0; i < len(s); i++ {
	// 	x := m[s[i]]

	// 	if i == len(s)-1 {
	// 		return result + x
	// 	}

	// 	y := m[s[i+1]]
	// 	if x < y {
	// 		result += y - x
	// 		i++
	// 	} else {
	// 		result += x
	// 	}
	// }
	// return result
}

var m = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

// @lc code=end

