/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 *
 * https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/description/
 *
 * algorithms
 * Hard (50.10%)
 * Likes:    170
 * Dislikes: 0
 * Total Accepted:    35.1K
 * Total Submissions: 69.9K
 * Testcase Example:  '[1,3,5]'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 *
 * ( 例如，数组 [0,1,2,4,5,6,7]  可能变为 [4,5,6,7,0,1,2] )。
 *
 * 请找出其中最小的元素。
 *
 * 注意数组中可能存在重复的元素。
 *
 * 示例 1：
 *
 * 输入: [1,3,5]
 * 输出: 1
 *
 * 示例 2：
 *
 * 输入: [2,2,2,0,1]
 * 输出: 0
 *
 * 说明：
 *
 *
 * 这道题是 寻找旋转排序数组中的最小值 的延伸题目。
 * 允许重复会影响算法的时间复杂度吗？会如何影响，为什么？
 *
 *
 */
package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMin_154(t *testing.T) {
	for index, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{3, 5, 1}, 1},
		{[]int{1, 3, 5}, 1},
		{[]int{2, 2, 2, 0, 1}, 0},
		{[]int{3, 1}, 1},
		{[]int{1, 1}, 1},
	} {
		t.Run(fmt.Sprintf("case: %v", index), func(t *testing.T) {
			assert.Equal(t, tc.want, findMin_154(tc.nums))
		})
	}
}

// @lc code=start
func findMin_154(nums []int) int {
	// func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)>>1

		// 旋转排序数组，用中间值与右侧值作比较
		// 与左侧值比较不行：
		// 例 1：[1, 2, 3, 4, 5]
		// 例 2：[2, 3, 4, 5, 1]
		if nums[mid] == nums[right] {
			right--
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return nums[left]
}

// func findMin(nums []int) int {
// 	min := nums[0]
// 	for _, num := range nums {
// 		if num < min {
// 			min = num
// 		}
// 	}
// 	return min
// }

// @lc code=end
