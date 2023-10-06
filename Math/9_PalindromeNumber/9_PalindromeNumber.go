package leetcode

// LeetCode：https://leetcode.com/problems/palindrome-number/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0009.Palindrome-Number/
// Difficulty：Easy
// Time Complexity：
// Space Complexity：

// note
// 判斷一個數字是否為回文(palindrome)從左到右，從右到左的順序都相同
// Follow up：能夠在不轉換成字串的情況下完成嗎？
// 筆記：從cookbook看來的解法，有優化減少一些if判斷跟不必要的cap宣告，較為精簡
// 20231006 複習一次解出來

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	arr := []int{}
	for x > 0 {
		arr = append(arr, x%10)
		x = x / 10
	}

	rightSide := len(arr) - 1
	for i, j := 0, rightSide; i <= j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}

	return true
}
