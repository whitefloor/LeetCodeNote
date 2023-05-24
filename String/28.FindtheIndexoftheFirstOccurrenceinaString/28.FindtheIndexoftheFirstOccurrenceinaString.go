package leetcode

// LeetCode：https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0028.Find-the-Index-of-the-First-Occurrence-in-a-String/
// Difficulty：Easy
// Time Complexity：
// Space Complexity：O(1)

//note
// 有參考解答
// 這題主要是在找出指定的子字串
// 有想過要在匹配的時候如果不匹配就直接往後移len(needle)，不過失敗

func strStr(haystack string, needle string) int {
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if j == len(needle) {
				return i
			}
			if i+j == len(haystack) {
				return -1
			}
			if haystack[i+j] != needle[j] {
				break
			}
		}
	}
}
