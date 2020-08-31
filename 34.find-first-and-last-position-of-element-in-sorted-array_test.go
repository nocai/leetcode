/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * <65;47;53M<64;40;44M<64;65;48M[34] 在排序数组中查找元素的第一个和最后一个位置
 *
 * https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (4018%)
 * Likes:    569
 * Dislikes: 0
 * Total Accepted:    125.3K
 * Total Submissions: 311.5K
 * Testcase Example:  '[<65;47;53M.5,7,7,8,8,10]\n8'
 *
 * 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
 *
 * 你的算法时间复杂度必须是 O(log ) 级别。
 *
 * 如果数组中不存在目标值，返回 [<65;47;53Mn-1, -1]。
 *
 * 示例 1:
 *
 * 输入: nums = [5,7,7,8,8,10], target = 8
 * 输出: [3,4]
 *
 * 示例 2:
 *
 * 输入: nums = [5,7,7,8,8,10], target = 6
 * 输出: [-1,-1]
 *
 */

package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func searchRange(nums []int, target int) []int {
	first := findFirstPosition(nums, target)
	if first == -1 {
		return []int{-1, -1}
	}

	last := findLastPosition(nums, target)
	return []int{first, last}
	// return []int{searchRangeLeft(nums, target), searchRangeRight(nums, target)}
}

func findLastPosition(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left+1)/2
		if nums[mid] > target {
			// [left, mid-1]
			right = mid - 1
		} else {
			// [mid, right]
			left = mid
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}

func findFirstPosition(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}

func searchRangeRight(nums []int, target int) int {
	// 特殊情况处理：
	// 1. nums为空
	if len(nums) == 0 {
		return -1
	}
	// 2. 不在范围内
	if target < nums[0] || nums[len(nums)-1] < target {
		return -1
	}

	left, right := 0, len(nums) // [left, right) => [left, mid) + mid + [mid+1, right)

	for left < right { // 终止条件：left == right
		mid := left + (right-left)>>1
		if target == nums[mid] {
			left = mid + 1 // 找右侧边界，左指针向右压缩范围
		} else if target < nums[mid] {
			right = mid
		} else if target > nums[mid] {
			left = mid + 1
		}
	}

	// 找到
	if nums[left-1] == target {
		return left - 1
	}
	return -1
}

func searchRangeLeft(nums []int, target int) int {
	// 特殊情况处理：
	// 1. nums为空
	if len(nums) == 0 {
		return -1
	}
	// 2. 不在范围内
	if target < nums[0] || nums[len(nums)-1] < target {
		return -1
	}

	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)>>1
		if target == nums[mid] {
			right = mid // 找左侧边界，右指针向左压缩范围
		} else if target < nums[mid] {
			right = mid
		} else if target > nums[mid] {
			left = mid + 1
		}
	}

	// 找到
	if nums[left] == target {
		return left
	}
	return -1
}

// @lc code=end

func TestSearchRangeLeft(t *testing.T) {
	for index, tc := range []struct {
		nums   []int
		target int
		index  int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, 3},
		{[]int{5, 7, 7, 8, 8, 10}, 0, -1},
		{[]int{5, 7, 7, 8, 8, 10}, 11, -1},
		{[]int{}, 11, -1},
		{[]int{5, 7, 7, 8, 8, 10}, 6, -1},
	} {
		t.Run(fmt.Sprintf("case: %v", index), func(t *testing.T) {
			assert.Equal(t, tc.index, searchRangeLeft(tc.nums, tc.target))
		})
	}
}

func TestSearchRangeRight(t *testing.T) {
	for index, tc := range []struct {
		nums   []int
		target int
		index  int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, 4},
		{[]int{5, 7, 7, 8, 8, 10}, 0, -1},
		{[]int{5, 7, 7, 8, 8, 10}, 11, -1},
		{[]int{}, 11, -1},
		{[]int{5, 7, 7, 8, 8, 10}, 6, -1},
	} {
		t.Run(fmt.Sprintf("case: %v", index), func(t *testing.T) {
			assert.Equal(t, tc.index, searchRangeRight(tc.nums, tc.target))
		})
	}
}

func TestSearchRange(t *testing.T) {
	for index, tc := range []struct {
		nums   []int
		target int
		rst    []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
	} {
		t.Run(fmt.Sprintf("case: %v", index), func(t *testing.T) {
			assert.Equal(t, tc.rst, searchRange(tc.nums, tc.target))
		})
	}
}
