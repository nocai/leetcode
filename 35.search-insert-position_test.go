/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 *
 * https://leetcode-cn.com/problems/search-insert-position/description/
 *
 * algorithms
 * Easy (46.64%)
 * Likes:    657
 * Dislikes: 0
 * Total Accepted:    235.1K
 * Total Submissions: 504.1K
 * Testcase Example:  '[1,3,5,6]\n5'
 *
 * 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
 *
 * 你可以假设数组中无重复元素。
 *
 * 示例 1:
 *
 * 输入: [1,3,5,6], 5
 * 输出: 2
 *
 *
 * 示例 2:
 *
 * 输入: [1,3,5,6], 2
 * 输出: 1
 *
 *
 * 示例 3:
 *
 * 输入: [1,3,5,6], 7
 * 输出: 4
 *
 *
 * 示例 4:
 *
 * 输入: [1,3,5,6], 0
 * 输出: 0
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
func searchInsert_2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	// 搜索区间：[left, right]
	for left <= right {
		mid := left + (right-left)>>1
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			// [left, mid-1]
			right = mid - 1
		} else if target > nums[mid] {
			// [mid+1, right]
			left = mid + 1
		}
	}
	return left
}

func searchInsert(nums []int, target int) int {
	// 初始值的取舍范围
	// 决定下面的搜索空间
	left, right := 0, len(nums)

	// 这里循环结束时，left == right
	// [left, right]
	for left < right {
		mid := left + (right-left)/2
		if target <= nums[mid] {
			// [left, mid]
			right = mid
		} else {
			// [mid+1, right]
			left = mid + 1
		}
		// if nums[mid] < target {
		// 	// [mid+1, right]
		// 	left = mid + 1
		// } else {
		// 	// [left, mid]
		// 	right = mid
		// }
	}
	return left
}

// @lc code=end

func TestSearchInsert(t *testing.T) {
	for index, tc := range []struct {
		nums   []int
		target int
		rst    int
	}{
		{[]int{1}, 1, 0},
		{[]int{2, 3, 5, 6}, 1, 0},
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
		{[]int{1, 3}, 2, 1},
	} {
		t.Run(fmt.Sprintf("case: %v", index), func(t *testing.T) {
			assert.Equal(t, tc.rst, searchInsert_2(tc.nums, tc.target))
		})
	}
}
