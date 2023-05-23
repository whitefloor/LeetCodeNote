package leetcode

// LeetCode：https://leetcode.com/problems/longest-common-prefix/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0014.Longest-Common-Prefix/
// Difficulty：Easy
// Time Complexity：O(minWordLen * words)
// Space Complexity：O(1)

//note
// 找出所有字串中共同的最長前綴
// 有參考解答，自己寫設太多條件寫有夠醜＝＝
// 只要先找到最短的字，之後找到共同前綴就加append，不一樣就可以break了
// 不過要注意的是有提示字串都是小寫英文字母，如果處理中文的話要注意會有UTF-8的問題

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	for i := 0; i < len(strs); i++ {
		for j := 0; j < len(prefix); j++ {
			if len(strs[i]) <= j || strs[i][j] != prefix[j] {
				prefix = prefix[0:j]
				break
			}
		}
	}

	return prefix
}
