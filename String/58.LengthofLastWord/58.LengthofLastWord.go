package leetcode

// LeetCode：https://leetcode.com/problems/length-of-last-word/solutions/3543731/go-solution/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0058.Length-of-Last-Word/?
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

//note
// 看了解答之後想出來用一個for迴圈就搞定的方法
// 解法從字串的最後開始找比較簡單
// 不過從solutions看到一個條件更簡單的
// https://leetcode.com/problems/length-of-last-word/solutions/3543731/go-solution/

// 自己寫的
func lengthOfLastWord(s string) int {
	start, end := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if start != 0 || i-1 < 0 {
			break
		}
		if end == 0 && s[i] != ' ' {
			end = i
		}
		if start == 0 && s[i] != ' ' && s[i-1] == ' ' {
			start = i
		}
	}

	return end - start + 1
}

// 條件更清晰的
func lengthOfLastWord(s string) int {
	var c int

	for i := len(s) - 1; ; i-- {
		if i < 0 {
			return c
		}

		if c > 0 && string(s[i]) == " " {
			return c
		}

		if string(s[i]) == " " {
			continue
		}

		c++
	}

	return 0
}
