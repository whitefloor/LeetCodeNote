package leetcode

// LeetCode：https://leetcode.com/problems/maximum-depth-of-binary-tree/description/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0104.Maximum-Depth-of-Binary-Tree/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

// note
// 經典的路徑查找，看哪裡的節點最深，總共有幾個節點
// 看解答寫出來的
// 特別注意 max() 是 leetcode 才有的特有功能，正常寫 code 是沒有這個 function的

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// leetcode 內建的，可以不用自己寫
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//
//	return b
//}
