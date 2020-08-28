/*
 * @lc app=leetcode.cn id=441 lang=golang
 *
 * [441] 排列硬币
 *
 * https://leetcode-cn.com/problems/arranging-coins/description/
 *
 * algorithms
 * Easy (41.21%)
 * Likes:    73
 * Dislikes: 0
 * Total Accepted:    27.4K
 * Total Submissions: 66.4K
 * Testcase Example:  '5'
 *
 * 你总共有 n 枚硬币，你需要将它们摆成一个阶梯形状，第 k 行就必须正好有 k 枚硬币。
 *
 * 给定一个数字 n，找出可形成完整阶梯行的总行数。
 *
 * n 是一个非负整数，并且在32位有符号整型的范围内。
 *
 * 示例 1:
 *
 *
 * n = 5
 *
 * 硬币可排列成以下几行:
 * ¤
 * ¤ ¤
 * ¤ ¤
 *
 * 因为第三行不完整，所以返回2.
 *
 *
 * 示例 2:
 *
 *
 * n = 8
 *
 * 硬币可排列成以下几行:
 * ¤
 * ¤ ¤
 * ¤ ¤ ¤
 * ¤ ¤
 *
 * 因为第四行不完整，所以返回3.
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

func arrangeCoins(n int) int {
	// 寻找右侧边界的二分查找
	left, right := 1, n+1
	for left < right {
		mid := left + (right-left)/2
		cost := ((mid + 1) * mid) / 2
		if cost == n {
			return mid
		} else if cost < n {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left - 1
}

// func arrangeCoins(n int) int {
// 	sum := 0
// 	for i := 0; i <= n; i++ {
// 		sum += i
// 		if sum <= n && n < sum+i+1 {
// 			return i
// 		}
// 	}
// 	return 0
// }

// @lc code=end
func TestArrangeCoins(t *testing.T) {
	for index, tc := range []struct {
		in  int
		out int
	}{
		{0, 0},
		{1, 1},
		{5, 2},
		{8, 3},
	} {
		t.Run(fmt.Sprintf("case: %v", index), func(t *testing.T) {
			assert.Equal(t, tc.out, arrangeCoins(tc.in))
		})
	}
}
