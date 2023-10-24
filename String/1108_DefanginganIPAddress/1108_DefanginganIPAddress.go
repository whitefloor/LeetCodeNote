package leetcode

// LeetCode：https://leetcode.com/problems/defanging-an-ip-address/description/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/1100~1199/1108.Defanging-an-IP-Address/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：

// note
// 給你一個 IPv4 的地址字串，將其中的 "." 替換成 "[.]"
// 簡單題，沒看解答
// 解答用的是 package

func defangIPaddr(address string) string {
	output := ""
	for _, v := range address {
		if findNum(v) {
			output += string(v)
		} else {
			output += "[.]"
		}
	}
	return string(output)
}

func findNum(s rune) bool {
	if '0' <= s && s <= '9' {
		return true
	}
	return false
}
