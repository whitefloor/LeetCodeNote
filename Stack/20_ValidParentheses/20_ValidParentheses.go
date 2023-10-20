package leetcode

// LeetCode：https://leetcode.com/problems/valid-parentheses/description/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0020.Valid-Parentheses/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

// note
// 給你一串由 [] {} () 三種括號組成的字串，判斷是否其兩兩成雙
// 看解答寫出來的，只要有匹配就把 stack 的頭給頂出去

func isValid(s string) bool {
	m := map[rune]rune{}
	stack := []rune{}
	m['('] = ')'
	m['{'] = '}'
	m['['] = ']'

	for _, v := range s {
		if _, ok := m[v]; ok {
			stack = append(stack, v)
		} else if len(stack) > 0 && m[stack[len(stack)-1]] == v {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}
