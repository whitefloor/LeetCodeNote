package leetcode

// LeetCode：https://leetcode.com/problems/pascals-triangle/description/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0118.Pascals-Triangle/
// Difficulty：Easy
// Time Complexity：
// Space Complexity：

// note
// 看解答想出來的，其實沒什麼難
// 主要是帕斯卡三角形的邏輯

func generate(numRows int) [][]int {
	res := [][]int{}

	for i := 0; i < numRows; i++ {
		arr := []int{}
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				arr = append(arr, 1)
			} else if i > 1 {
				arr = append(arr, res[i-1][j]+res[i-1][j-1])
			}
		}
		res = append(res, arr)
	}

	return res
}
