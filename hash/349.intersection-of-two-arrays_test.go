/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 *
 * https://leetcode-cn.com/problems/intersection-of-two-arrays/description/
 *
 * algorithms
 * Easy (70.65%)
 * Total Accepted:    95K
 * Total Submissions: 134.2K
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
func intersection(nums1 []int, nums2 []int) []int {
	s1 := map[int]struct{}{}
	for _, num := range nums1 {
		s1[num] = struct{}{}
	}

	s2 := map[int]struct{}{}
	for _, num := range nums2 {
		s2[num] = struct{}{}
	}

	m, l := s1, s2
	if len(s1) > len(s2) {
		m, l = s2, s1
	}

	rst := make([]int, 0, len(m))
	for v, _ := range m {
		if _, ex := l[v]; ex {
			rst = append(rst, v)
		}
	}
	return rst
}
