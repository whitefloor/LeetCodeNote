package leetcode

// LeetCode：https://leetcode.com/problems/pascals-triangle-ii/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0119.Pascals-Triangle-II/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(rowIndex)

// note
// 118的進階版本，需要把 space complexity 優化到 O(rowIndex)
// 基本上要知道帕斯卡三角形的公式才能解

func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 1; i <= rowIndex; i++ {
		row[i] = row[i-1] * (rowIndex - i + 1) / i
	}

	return row
}
