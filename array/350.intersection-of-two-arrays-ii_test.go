/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 *
 * https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/description/
 *
 * algorithms
 * Easy (52.47%)
 * Likes:    381
 * Dislikes: 0
 * Total Accepted:    136.6K
 * Total Submissions: 259.8K
 * Testcase Example:  '[1,2,2,1]\n[2,2]'
 *
 * 给定两个数组，编写一个函数来计算它们的交集。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums1 = [1,2,2,1], nums2 = [2,2]
 * 输出：[2,2]
 *
 *
 * 示例 2:
 *
 * 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * 输出：[4,9]
 *
 *
 *
 * 说明：
 *
 *
 * 输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
 * 我们可以不考虑输出结果的顺序。
 *
 *
 * 进阶：
 *
 *
 * 如果给定的数组已经排好序呢？你将如何优化你的算法？
 * 如果 nums1 的大小比 nums2 小很多，哪种方法更优？
 * 如果 nums2 的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
 *
 *
 */
package leetcode

import "sort"

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int { // O(n)
	// 先排序后，再双指针

	sort.Ints(nums1)
	sort.Ints(nums2)

	p1, p2 := 0, 0
	result := []int{}
	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] == nums2[p2] {
			result = append(result, nums1[p1])
			p1++
			p2++
		} else if nums1[p1] < nums2[p2] {
			p1++
		} else {
			p2++
		}
	}

	return result
}

// func intersect(nums1 []int, nums2 []int) []int { // O(m+n)
// 	m := map[int]int{}
// 	for _, v := range nums1 {
// 		m[v]++
// 	}

// 	result := []int{}
// 	for _, v := range nums2 {
// 		if m[v] > 0 {
// 			m[v]--
// 			result = append(result, v)
// 		}
// 	}
// 	return result
// }

// @lc code=end
