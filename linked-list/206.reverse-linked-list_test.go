/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 *
 * https://leetcode-cn.com/problems/reverse-linked-list/description/
 *
 * algorithms
 * Easy (70.28%)
 * Likes:    1216
 * Dislikes: 0
 * Total Accepted:    329K
 * Total Submissions: 466.6K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 反转一个单链表。
 *
 * 示例:
 *
 * 输入: 1->2->3->4->5->NULL
 * 输出: 5->4->3->2->1->NULL
 *
 * 进阶:
 * 你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var (
		prev *ListNode = nil
		cur            = head
	)
	for cur != nil {
		// 用一个临时节点保存当前节点的下一节点
		tmp := cur.Next
		// 指向下一节点的指针反向
		cur.Next = prev
		// 两个节点向右移
		prev = cur
		cur = tmp
	}
	return prev
}

func reverseList2(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head

	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prve = cur
		cur = tmp
	}
	return prev
}

// @lc code=end
