package leetcode

// LeetCode：https://leetcode.com/problems/reverse-string/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0300~0399/0344.Reverse-String/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

// note
// 使用 O(1) 的空間，把字串反轉
// 蠻簡單的，一次就解出來，只要頭尾不斷交換就可以

func reverseString(s []byte) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
