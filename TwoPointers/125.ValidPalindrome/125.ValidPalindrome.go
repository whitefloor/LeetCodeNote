package leetcode

import "strings"

// LeetCode：https://leetcode.com/problems/valid-palindrome/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0125.Valid-Palindrome/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

// note
// palindrome:回文的意思，意思是說一段字串從前看跟從後看都會一模一樣
// 這題有點爛，因為要把字變小寫，忘了 API 就很難處理
// 從字串呼叫個別 string 型別是 int8 = bytes
// 從 for range 呼叫 string 型別是 int32 = rune
// 'a' 的方式呼叫是 rune

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isChar(s[i]) {
			i++
		}
		for i < j && !isChar(s[j]) {
			j--
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func isChar(s byte) bool {
	if ('a' <= s && s <= 'z') || ('0' <= s && s <= '9') {
		return true
	}
	return false
}
