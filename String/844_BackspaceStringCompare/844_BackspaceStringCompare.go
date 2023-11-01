package leetcode

// LeetCode：https://leetcode.com/problems/backspace-string-compare/description/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0800~0899/0844.Backspace-String-Compare/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：

// note
// 給定兩個字串 s,t 比較這兩個字串是否相等，中間會包含空白部分，用 # 代替
// 字串遇到 # 會回退，例如 ab#c = ac
// 看解答寫出來的，利用的是用 stack 的方法，後進先出
// 一開始想法蠻接近的，但是沒寫出來

func backspaceCompare(s string, t string) bool {
	var ss, ts []rune
	for _, v := range s {
		if v == '#' {
			if len(ss) > 0 {
				ss = ss[:len(ss)-1]
			}
		} else {
			ss = append(ss, v)
		}
	}
	for _, v := range t {
		if v == '#' {
			if len(ts) > 0 {
				ts = ts[:len(ts)-1]
			}
		} else {
			ts = append(ts, v)
		}
	}
	return string(ss) == string(ts)
}
