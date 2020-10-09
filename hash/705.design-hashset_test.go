/*
 * @lc app=leetcode.cn id=705 lang=golang
 *
 * [705] Design HashSet
 *
 * https://leetcode-cn.com/problems/design-hashset/description/
 *
 * algorithms
 * Easy (57.49%)
 * Total Accepted:    18.7K
 * Total Submissions: 32.6K
 * Testcase Example:  '["MyHashSet","add","add","contains","contains","add","contains","remove","contains"]\n' +
  '[[],[1],[2],[1],[3],[2],[2],[2],[2]]'
 *
 * 不使用任何内建的哈希表库设计一个哈希集合
 *
 * 具体地说，你的设计应该包含以下的功能
 *
 *
 * add(value)：向哈希集合中插入一个值。
 * contains(value) ：返回哈希集合中是否存在这个值。
 * remove(value)：将给定值从哈希集合中删除。如果哈希集合中没有这个值，什么也不做。
 *
 *
 *
 * 示例:
 *
 * MyHashSet hashSet = new MyHashSet();
 * hashSet.add(1);
 * hashSet.add(2);
 * hashSet.contains(1);    // 返回 true
 * hashSet.contains(3);    // 返回 false (未找到)
 * hashSet.add(2);
 * hashSet.contains(2);    // 返回 true
 * hashSet.remove(2);
 * hashSet.contains(2);    // 返回  false (已经被删除)
 *
 *
 *
 * 注意：
 *
 *
 * 所有的值都在 [0, 1000000]的范围内。
 * 操作的总数目在[1, 10000]范围内。
 * 不要使用内建的哈希集合库。
 *
 *
*/
type MyHashSet struct {
	Buckets []list.List
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{
		Buckets: make([]list.List, 0, 1000),
	}
}

func (this *MyHashSet) hash(key int) int {
	return key % len(this.Buckets)
}

func (this *MyHashSet) Add(key int) {
	index := this.hash(key)
	this.Buckets[index].PushBack(key)
}

func (this *MyHashSet) Remove(key int) {
	index := this.hash(key)

	l := this.Buckets[index]
	for ele := l.Front(); ele != nil; ele = ele.Next() {
		if ele.Value == key {
			l.Remove(ele)
			return
		}
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	index := this.hash(key)
	ele := this.Buckets[index].Front()
	for ele != nil {
		if ele.Value == key {
			return true
		}
		ele = ele.Next()
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

