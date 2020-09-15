package leetcode

import "testing"

func TestList(t *testing.T) {
	// 一。在链表中间添加节点
	// 1. cur := new(cur)
	// 2. cur.Next = prev.Next
	// 3. prev.Next = cur

	// 二。在链表头添加节点
	// 1. cur := new(cur) // 初始化一个新节点
	// 2. cur.Next = head // 将新节点链接到我们原始的头结点head上
	// 3. head = cur // 将cur指定为head

	// 三。在链表中间删除节点
	// 1. 找到prev, cur, next
	// 2. prev.Next = cur.Next

	// 四。在链表头删除
	// 1. 简单地“将下一个结点分配给head” head = head.Next
}

func TestList2(t *testing.T) {
	// 双向链表操作

	// 一。添加
	// 1. 链接cur与prev和next, 其中next是prev原始的下一个节点
	// cur.Prev = prev, cur.Next = next
	// 2. 用cur重新链接prev和next
	// prev.Next = cur, next.Prev = cur

	// 二。删除
	// 如果我们想从双链表中删除一个现有的结点 cur，我们可以简单地将它的前一个结点 prev 与下一个结点 next 链接起来。
	// cur.Prev.Next = cur.Next, cur.Next.Prev = cur.Prev
}

type MyLinkedList struct{}

func (this *MyLinkedList) Get(index int) int {
	return 0
}

func (this *MyLinkedList) AddAtHead(val int) {

}

func (this *MyLinkedList) AddAdTail(val int) {

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {

}

func (this *MyLinkedList) DeleteAdIndex(index int) {

}

// Definition for singly-linked list.
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func (ln *ListNode) String() string {
// 	var rst strings.Builder
// 	rst.WriteString("[")
// 	for ln != nil {
// 		rst.WriteString(fmt.Sprintf("%v", ln.Val))
// 		if ln.Next != nil {
// 			rst.WriteRune(',')
// 		}
// 		ln = ln.Next
// 	}
// 	rst.WriteString("]")
// 	return rst.String()
// }

// func InitLinkedList(vals ...int) *ListNode {
// 	if len(vals) == 0 {
// 		return nil
// 	}

// 	head := &ListNode{Val: vals[0]}
// 	tail := head

// 	for _, val := range vals[1:] {
// 		cur := &ListNode{Val: val}
// 		tail.Next = cur
// 		tail = tail.Next
// 	}
// 	return head
// }

// func TestInitLinkedList(t *testing.T) {
// 	t.Log(InitLinkedList(1, 2, 3))
// }

// // func hasCycle(head *ListNode) bool {

// // 	return false
// // }

// func TestHasCycle_hash(t *testing.T) {
// 	for index, tc := range []struct {
// 		head     *ListNode
// 		hasCycle bool
// 	}{

// 		{InitLinkedList(3, 2, 0, -4), true},
// 		{InitLinkedList(1, 2), true},
// 		{InitLinkedList(1), false},
// 	} {
// 		t.Run(fmt.Sprintf("case: %v", index), func(t *testing.T) {
// 			// assert.Equal(t, tc.hasCycle, hasCycle_hash(tc.head))
// 		})
// 	}
// }
