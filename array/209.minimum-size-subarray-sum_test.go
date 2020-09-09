/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 *
 * https://leetcode-cn.com/problems/minimum-size-subarray-sum/description/
 *
 * algorithms
 * Medium (44.36%)
 * Likes:    429
 * Dislikes: 0
 * Total Accepted:    86.2K
 * Total Submissions: 194.3K
 * Testcase Example:  '7\n[2,3,1,2,4,3]'
 *
 * 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续
 * 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。
 *
 *
 *
 * 示例：
 *
 * 输入：s = 7, nums = [2,3,1,2,4,3]
 * 输出：2
 * 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 如果你已经完成了 O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
 *
 *
 */

package leetcode

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func minSubArrayLen(s int, nums []int) int {
	return 0
}

// @lc code=end

func TestMinSubArrayLen(t *testing.T) {
	for index, tc := range []struct {
		nums []int
		s    int
		want int
	}{
		{[]int{2, 3, 1, 2, 4, 3}, 7, 2},
	} {
		t.Run(fmt.Sprintf("case: %v", index), func(t *testing.T) {
			assert.Equal(t, tc.want, minSubArrayLen_two_pointers(tc.s, tc.nums))
		})
	}
}

func minSubArrayLen_prefixSum_binary_search(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	rst := math.MaxInt64
	sums := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			sums[0] = nums[0]
		} else {
			sums[i] = sums[i-1] + nums[i]
		}
	}

	// TODO: 后面不会写了 <08-09-20, lj.liujun> //
	_ = rst
	return 0
}

func minSubArrayLen_two_pointers(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var (
		rst        = math.MaxInt64
		slow, fast = 0, 0
		sum        = 0
	)

	for fast < len(nums) {
		sum += nums[fast]
		for sum >= s {
			rst = min(rst, fast-slow+1)
			sum -= nums[slow]
			slow++
		}
		fast++
	}

	if rst == math.MaxInt64 {
		return 0
	}
	return rst
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minSubArrayLen_brute(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	rst := math.MaxInt64
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum >= s {
				l := j - i + 1
				if l < rst {
					rst = l
				}
				break
			}
		}
	}

	if rst == math.MaxInt64 {
		return 0
	}
	return rst
}
