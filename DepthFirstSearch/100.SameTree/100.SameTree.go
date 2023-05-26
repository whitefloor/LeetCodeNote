package leetcode

// LeetCode：https://leetcode.com/problems/same-tree/submissions/957548305/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0100.Same-Tree/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

//note
// 看解答的，對DFS跟Tree還不是很熟
// 不過這題說真的不需要太多的操作，只需要判斷整棵樹的node跟value是不是相同即可

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

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
