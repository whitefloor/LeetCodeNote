package leetcode

import "strconv"

// LeetCode：https://leetcode.com/problems/binary-tree-paths/description/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0200~0299/0257.Binary-Tree-Paths/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

// note
// 給一個 Tree，不限順序的 traversal 輸出，但要加上每一種路徑以及符號，最後用字串陣列輸出
// 看解答寫出來的，中序遍歷沒問題（中、左、右）

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []string

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	res = []string{}
	traversalNode(root, "")

	return res
}

func traversalNode(root *TreeNode, path string) string {
	if path == "" {
		path = strconv.Itoa(root.Val)
	} else {
		path = path + "->" + strconv.Itoa(root.Val)
	}

	if root.Left != nil {
		traversalNode(root.Left, path)
	}
	if root.Right != nil {
		traversalNode(root.Right, path)
	}

	if root.Left == nil && root.Right == nil {
		res = append(res, path)
	}

	return path
}
