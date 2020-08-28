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

func TestFindMin2(t *testing.T) {
	for index, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 3, 5}, 1},
		{[]int{2, 2, 2, 0, 1}, 0},
		{[]int{3, 1}, 1},
		{[]int{1, 1}, 1},
	} {
		t.Run(fmt.Sprintf("case: %v", index), func(t *testing.T) {
			assert.Equal(t, tc.want, findMin2(tc.nums))
		})
	}
}

// @lc code=start
// 我们考虑数组中的最后一个元素 xx：在最小值右侧的元素，它们的值一定都小于等于 xx；而在最小值左侧的元素，它们的值一定都大于等于 xx。因此，我们可以根据这一条性质，通过二分查找的方法找出最小值
func findMin2(nums []int) int {
	left, right := 0, len(nums)-1

	// 在二分查找的每一步中，左边界为 \it lowlow，右边界为 \it highhigh，区间的中点为 \it pivotpivot，最小值就在该区间内。我们将中轴元素 \textit{numbers}[\textit{pivot}]numbers[pivot] 与右边界元素 \textit{numbers}[\textit{high}]numbers[high] 进行比较，可能会有以下的三种情况：

	for left <= right {
		mid := left + (right-left)>>1

		if nums[mid] < nums[right] {
			// 第一种情况是 \textit{numbers}[\textit{pivot}] < \textit{numbers}[\textit{high}]numbers[pivot]<numbers[high]。如下图所示，这说明 \textit{numbers}[\textit{pivot}]numbers[pivot] 是最小值右侧的元素，因此我们可以忽略二分查找区间的右半部分。
			right = mid
		} else if nums[mid] > nums[right] {
			// 第二种情况是 \textit{numbers}[\textit{pivot}] > \textit{numbers}[\textit{high}]numbers[pivot]>numbers[high]。如下图所示，这说明 \textit{numbers}[\textit{pivot}]numbers[pivot] 是最小值左侧的元素，因此我们可以忽略二分查找区间的左半部分。
			left = mid + 1
		} else {
			// 第三种情况是 \textit{numbers}[\textit{pivot}] == \textit{numbers}[\textit{high}]numbers[pivot]==numbers[high]。如下图所示，由于重复元素的存在，我们并不能确定 \textit{numbers}[\textit{pivot}]numbers[pivot] 究竟在最小值的左侧还是右侧，因此我们不能莽撞地忽略某一部分的元素。我们唯一可以知道的是，由于它们的值相同，所以无论 \textit{numbers}[\textit{high}]numbers[high] 是不是最小值，都有一个它的「替代品」\textit{numbers}[\textit{pivot}]numbers[pivot]，因此我们可以忽略二分查找区间的右端点。
			right--
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
