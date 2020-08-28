/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 *
 * https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/description/
 *
 * algorithms
 * Easy (56.54%)
 * Likes:    386
 * Dislikes: 0
 * Total Accepted:    149.7K
 * Total Submissions: 264.8K
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * 给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
 *
 * 函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。
 *
 * 说明:
 *
 *
 * 返回的下标值（index1 和 index2）不是从零开始的。
 * 你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
 *
 *
 * 示例:
 *
 * 输入: numbers = [2, 7, 11, 15], target = 9
 * 输出: [1,2]
 * 解释: 2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。
 *
 */

package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum2(t *testing.T) {
	for index, tc := range []struct {
		numbers []int
		target  int
		rst     []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 2, 11, 15}, 4, []int{1, 2}},
		{[]int{2, 2, 11, 15}, -1, []int{-1, -1}},
	} {
		t.Run(fmt.Sprintf("case: %v", index), func(t *testing.T) {
			assert.Equal(t, tc.rst, twoSum2(tc.numbers, tc.target))
		})
		t.Run(fmt.Sprintf("twoSum_binary_search: %v", index), func(t *testing.T) {
			assert.Equal(t, tc.rst, twoSum_binary_search(tc.numbers, tc.target))
		})
		t.Run(fmt.Sprintf("twoSum_hash: %v", index), func(t *testing.T) {
			assert.Equal(t, tc.rst, twoSum_hash(tc.numbers, tc.target))
		})
	}
}

// @lc code=start
func twoSum2(numbers []int, target int) []int {
	// double points

	left, right := 0, len(numbers)-1 // [0, len(numbers)-1]
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}

		if sum < target {
			left++
		} else if sum > target {
			right--
		}
	}
	return []int{-1, -1}
}

func twoSum_binary_search(numbers []int, target int) []int {
	// binary search

	for index := range numbers {
		left, right := index+1, len(numbers)-1 // [index+1, len(numbers)-1]

		_target := target - numbers[index]
		for left <= right {
			mid := left + (right-left)>>1
			if numbers[mid] == _target {
				return []int{index + 1, mid + 1}
			}

			if numbers[mid] < _target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return []int{-1, -1}
}

func twoSum_hash(numbers []int, target int) []int {
	// func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)
	for index, num := range numbers {
		key := target - num
		if idx, exist := m[key]; exist {
			// 较小的index放前面
			if index > idx {
				index, idx = idx, index
			}
			return []int{index + 1, idx + 1}
		} else {
			m[num] = index
		}
	}
	return []int{-1, -1}
}

// @lc code=end
