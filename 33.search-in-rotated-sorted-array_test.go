/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 *
 * https://leetcode-cn.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (38.54%)
 * Likes:    911
 * Dislikes: 0
 * Total Accepted:    163K
 * Total Submissions: 421.6K
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 *
 * ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
 *
 * 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
 *
 * 你可以假设数组中不存在重复的元素。
 *
 * 你的算法时间复杂度必须是 O(log n) 级别。
 *
 * 示例 1:
 *
 * 输入: nums = [4,5,6,7,0,1,2], target = 0
 * 输出: 4
 *
 *
 * 示例 2:
 *
 * 输入: nums = [4,5,6,7,0,1,2], target = 3
 * 输出: -1
 *
 */
package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func search_brute_force(nums []int, target int) int {
	for index, num := range nums {
		if num == target {
			return index
		}
	}
	return -1
}

func search_33_binary_search_2(nums []int, target int) int {
	// func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[left] {
			// 左边[left,mid]有序
			// 判断taget在此范围没有
			if nums[left] <= target && target <= nums[mid] {
				// 在
				// 下一轮范围：[left,mid)
				right = mid - 1
			} else {
				// (mid, right]
				left = mid + 1
			}
		} else {
			// 右边(mid, right]有序
			if nums[mid] < target && target <= nums[right] {
				// (mid, right]
				left = mid + 1
			} else {
				// [left, mid)
				right = mid - 1
			}
		}
	}
	return -1
}

func search_33_binary_search(nums []int, target int) int {
	// func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= nums[left] {
			// 左边[left,mid]有序
			// 判断taget在此范围没有
			if nums[left] <= target && target <= nums[mid] {
				// 在
				// 下一轮范围：[left,mid]
				right = mid
			} else {
				// (mid, right]
				left = mid + 1
			}
		} else {
			// 右边有序
			// (mid, right]
			if nums[mid] < target && target <= nums[right] {
				// (mid, right]
				left = mid + 1
			} else {
				// [left, mid]
				right = mid
			}
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}
func search2(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	firstVal, lastVal := nums[left], nums[right]

	for left <= right {
		mid := left + (right-left)>>1

		if target == nums[mid] {
			return mid
		}

		if firstVal <= nums[mid] {
			// 前半部分有序
			if firstVal <= target && target < nums[mid] {
				// 在前半部分范围内，right左移
				right = mid - 1
			} else {
				// 不在
				left = mid + 1
			}
		} else if firstVal > nums[mid] {
			// 后半部分有序
			if nums[mid] < target && target <= lastVal {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

// @lc code=end

func TestSearch2(t *testing.T) {
	for index, tc := range []struct {
		nums   []int
		target int
		index  int
	}{
		{[]int{3, 1}, 1, 1},
		{[]int{1}, 1, 0},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
	} {
		t.Run(fmt.Sprintf("case: %v", index), func(t *testing.T) {
			assert.Equal(t, tc.index, search_33_binary_search(tc.nums, tc.target))
		})
		t.Run(fmt.Sprintf("case_2: %v", index), func(t *testing.T) {
			assert.Equal(t, tc.index, search_33_binary_search_2(tc.nums, tc.target))
		})
	}
}
