package leetcode

// LeetCode：https://leetcode.com/problems/reverse-string/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0300~0399/0344.Reverse-String/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

//note
// 給你一個字串陣列反轉
// 只要頭尾一直持續反轉即可

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
