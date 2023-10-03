package leetcode

// LeetCode：https://leetcode.com/problems/happy-number/description/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0200~0299/0202.Happy-Number/
// Difficulty：Easy
// Time Complexity：
// Space Complexity：

// note
// 找出個快樂數字，在一開始給定的 int 將每個位數的平方和相加後為新的數字
// 只要在相加的過程中發現有不為1的循環就不是快樂數字
// 最後數字的每個位數平方和為1即為快樂數字
// 20231003 有參考解答，但有改進程式碼

func isHappy(n int) bool {
	record := map[int]bool{}
	for n != 1 {
		record[n] = true
		n = getSquareOfDigits(n)
		if _, ok := record[n]; ok {
			return false
		}
	}
	return true
}

func getSquareOfDigits(n int) int {
	squareOfDigits := 0
	for n != 0 {
		temp := n % 10
		squareOfDigits += temp * temp
		n = n / 10
	}
	return squareOfDigits
}
