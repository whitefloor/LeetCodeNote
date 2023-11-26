package leetcode

// LeetCode：https://leetcode.com/problems/valid-anagram/description/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0200~0299/0242.Valid-Anagram/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：

// note
// 题目意思是给你單詞，然後重新排列後成為字謎，判斷字謎是不是原本的那個單詞
// 題目蠻簡單的，用 map 判斷即可
// 不過在測試案例上少想了，導致解答不正確，有參考解答

func isAnagram(s string, t string) bool {
	m := map[rune]int{}
	for _, v := range s {
		m[v]++
	}

	for _, v := range t {
		m[v]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}
