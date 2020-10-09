/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 *
 * https://leetcode-cn.com/problems/first-unique-character-in-a-string/description/
 *
 * algorithms
 * Easy (46.92%)
 * Total Accepted:    104.1K
 * Total Submissions: 220.2K
 * Testcase Example:  '"leetcode"'
 *
 * 给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
 * 
 * 
 * 
 * 示例：
 * 
 * s = "leetcode"
 * 返回 0
 * 
 * s = "loveleetcode"
 * 返回 2
 * 
 * 
 * 
 * 
 * 提示：你可以假定该字符串只包含小写字母。
 * 
 */
func firstUniqChar(s string) int {
  h := map[byte]int{}

  for i := 0; i < len(s); i++ {
    h[s[i]]++
  }

  key byte := 0
  for k, v := range h {
    if v == 1 {
      key = k
      break
    }
  }

  for i := 0; i < len(s); i++ {
    if s[i] == key {
      return i
    }
  }
  return -1
}
