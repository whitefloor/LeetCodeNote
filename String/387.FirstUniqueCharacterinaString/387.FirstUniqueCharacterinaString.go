package leetcode

// LeetCode：https://leetcode.com/problems/first-unique-character-in-a-string/description/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0300~0399/0387.First-Unique-Character-in-a-String/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

//note
// 解法的話只要統計每個字母是不是有重複
// 之後在掃描一次字串看哪個字母的總數是1就可以return
// 不過網路上有的解法很厲害，利用ASCII當陣列index，最後value存數值

func firstUniqChar(s string) int {
	m := make(map[string]int)
	for _, str := range s {
		m[string(str)]++
	}

	for i, str := range s {
		if total := m[string(str)]; total == 1 {
			return i
		}
	}

	return -1
}
