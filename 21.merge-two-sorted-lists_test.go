/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 *
 * https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (63.97%)
 * Likes:    1224
 * Dislikes: 0
 * Total Accepted:    346.9K
 * Total Submissions: 542.1K
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
 *
 *
 *
 * 示例：
 *
 * 输入：1->2->4, 1->3->4
 * 输出：1->1->2->3->4->4
 *
 *
 */

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLosts(t *testing.T) {
	l1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	l2 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	got := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
		},
	}

	assert.Equal(t, &got, mergeTwoLists(&l1, &l2))
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 声明一个头、尾
	// 开始指向同一个地址
	var head ListNode
	tail := &head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			// tail指向l1
			tail.Next = l1
			// l1指针后移一个位置
			l1 = l1.Next
		} else {
			// tail指向l2
			tail.Next = l2
			// l2指针后移一个位置
			l2 = l2.Next
		}

		// tail指针后移一个位置
		tail = tail.Next
	}

	// 将l1/l2剩余的节点拼接到tail上
	if l1 != nil {
		tail.Next = l1
	} else if l2 != nil {
		tail.Next = l2
	}

	return head.Next

	// if l1 == nil {
	// 	return l2
	// }

	// if l2 == nil {
	// 	return l1
	// }

	// if l1.Val < l2.Val {
	// 	l1.Next = mergeTwoLists(l1.Next, l2)
	// 	return l1
	// } else {
	// 	l2.Next = mergeTwoLists(l2.Next, l1)
	// 	return l2
	// }
}

// @lc code=end
