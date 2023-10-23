package leetcode

// LeetCode：https://leetcode.com/problems/binary-tree-preorder-traversal/description/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0144.Binary-Tree-Preorder-Traversal/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：

// note
// preorder traversal = 前序遍歷
// 順序是 根>左>右，簡單題，秒解

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	output := []int{}
	preorder(root, &output)
	return output
}

func preorder(root *TreeNode, output *[]int) {
	if root != nil {
		*output = append(*output, root.Val)
		preorder(root.Left, output)
		preorder(root.Right, output)
	}
}
