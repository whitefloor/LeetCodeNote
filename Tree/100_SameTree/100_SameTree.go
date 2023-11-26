package leetcode

// LeetCode：https://leetcode.com/problems/same-tree/submissions/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0100.Same-Tree/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

// note
// 用遞迴在樹中做比較
// 遞迴跟部分比較的條件有寫出來，一部分參考解答

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else {
		return false
	}
}
