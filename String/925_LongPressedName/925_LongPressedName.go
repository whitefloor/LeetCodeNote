package leetcode

// LeetCode：https://leetcode.com/problems/long-pressed-name/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0900~0999/0925.Long-Pressed-Name/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

// note
// 給你 name 和 typed 來判斷 typed 是否符合 name，只是具有重複輸入的字元
// 看解答寫出來的，最後的條件式要記一下，有時候很好用
// 用滑動窗口的解法解出來的

func isLongPressedName(name string, typed string) bool {
	i, j := 0, 0
	for i < len(name) && j < len(typed) {
		if name[i] != typed[j] {
			return false
		}
		for i < len(name) && j < len(typed) && name[i] == typed[j] {
			i++
			j++
		}
		for j < len(typed) && typed[j] == typed[j-1] {
			j++
		}
	}

	return i == len(name) && j == len(typed)
}
