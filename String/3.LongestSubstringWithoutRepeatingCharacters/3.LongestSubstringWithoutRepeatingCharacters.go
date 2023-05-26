package leetcode

// LeetCode：https://leetcode.com/problems/longest-substring-without-repeating-characters/submissions/956991809/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0003.Longest-Substring-Without-Repeating-Characters/
// Difficulty：Medium
// Time Complexity：O(n)
// Space Complexity：O(1)

//note
// 想法跟cookbook解法三是一樣的，不過沒解出來有看解答
// 自己也有改寫code，自己比較看得懂
// 用map的特性做key-value處理重複的字元改變index蠻厲害的

func lengthOfLongestSubstring(s string) int {
	left, right, longest := 0, 0, 0
	m := map[byte]int{}
	for right < len(s) {
		if r, ok := m[s[right]]; ok && r >= left {
			left = r + 1
		}
		m[s[right]] = right
		right++

		long := right - left
		if long > longest {
			longest = long
		}
	}

	return longest
}
