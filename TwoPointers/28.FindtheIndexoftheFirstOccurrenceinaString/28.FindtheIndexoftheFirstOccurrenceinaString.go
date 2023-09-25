package leetcode

// LeetCode：https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0028.Find-the-Index-of-the-First-Occurrence-in-a-String/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

// note
// 在母字串中找到子字串並返回第一次找到相符子字串在母字串最前面的index
// 20230925 差點解出來，少想到了一些條件，最後看解答

func strStr(haystack string, needle string) int {
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if j == len(needle) {
				return i
			}
			if i+j == len(haystack) {
				return -1
			}
			if needle[j] != haystack[i+j] {
				break
			}
		}
	}
	// return -1 leetcode require this line but IDE show alert so remark
}
