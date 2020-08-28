/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 *
 * https://leetcode-cn.com/problems/intersection-of-two-arrays/description/
 *
 * algorithms
 * Easy (70.37%)
 * Likes:    224
 * Dislikes: 0
 * Total Accepted:    89.1K
 * Total Submissions: 126.5K
 * Testcase Example:  '[1,2,2,1]\n[2,2]'
 *
 * 给定两个数组，编写一个函数来计算它们的交集。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums1 = [1,2,2,1], nums2 = [2,2]
 * 输出：[2]
 *
 *
 * 示例 2：
 *
 * 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * 输出：[9,4]
 *
 *
 *
 * 说明：
 *
 *
 * 输出结果中的每个元素一定是唯一的。
 * 我们可以不考虑输出结果的顺序。
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
func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{})
	for i := 0; i < len(nums1); i++ {
		set[nums1[i]] = struct{}{}
	}

	set2 := make(map[int]struct{})
	for _, num := range nums2 {
		if _, exist := set[num]; exist {
			set2[num] = struct{}{}
		}
	}

	rst := []int{}
	for key := range set2 {
		rst = append(rst, key)
	}

	return rst
}

// @lc code=end

func TestIntersection(t *testing.T) {
	for index, tc := range []struct {
		nums1 []int
		nums2 []int
		rst   []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{9, 4}},
	} {
		t.Run(fmt.Sprintf("case: %v", index), func(t *testing.T) {
			assert.Equal(t, tc.rst, intersection(tc.nums1, tc.nums2))
		})
	}
}
