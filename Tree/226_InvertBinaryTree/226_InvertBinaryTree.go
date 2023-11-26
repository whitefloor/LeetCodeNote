package leetcode

// LeetCode：https://leetcode.com/problems/invert-binary-tree/description/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0200~0299/0226.Invert-Binary-Tree/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

// note
// 給你一個 binary tree 將他反轉，經典題
// 看解答才會的

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left

	return root
}
