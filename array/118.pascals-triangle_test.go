/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 *
 * https://leetcode-cn.com/problems/pascals-triangle/description/
 *
 * algorithms
 * Easy (67.15%)
 * Likes:    343
 * Dislikes: 0
 * Total Accepted:    103.4K
 * Total Submissions: 153.6K
 * Testcase Example:  '5'
 *
 * 给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
 *
 *
 *
 * 在杨辉三角中，每个数是它左上方和右上方的数的和。
 *
 * 示例:
 *
 * 输入: 5
 * 输出:
 * [
 * ⁠    [1],
 * ⁠   [1,1],
 * ⁠  [1,2,1],
 * ⁠ [1,3,3,1],
 * ⁠[1,4,6,4,1]
 * ]
 *
 */
package leetcode

import "testing"

// @lc code=start
func generate(numRows int) [][]int {
	rst := [][]int{}

	if numRows == 0 {
		return rst
	}

	rst = append(rst, []int{1})

	for i := 1; i < numRows; i++ {
		var (
			cur = make([]int, i)
			pre = rst[i-1]
		)

		cur[0] = 1
		for j := 1; j < i; j++ {
			cur[j] = pre[j-1] + pre[j]
		}
		cur = append(cur, 1)

		rst = append(rst, cur)
	}

	return rst
}

// @lc code=end

func TestGenrate(t *testing.T) {
	t.Logf("%v", generate(4))
}
