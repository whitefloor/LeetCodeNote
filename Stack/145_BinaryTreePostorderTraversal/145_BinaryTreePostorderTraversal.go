package leetcode

// LeetCode：https://leetcode.com/problems/binary-tree-postorder-traversal/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0145.Binary-Tree-Postorder-Traversal/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：

// note
// postorder traversal = 後序遍歷
// 順序是 左>右>根
// follow up 要求 使用 Recursive（遞迴)
// 簡單，一次過

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	output := []int{}
	postorder(root, &output)
	return output
}

func postorder(root *TreeNode, output *[]int) {
	if root != nil {
		postorder(root.Left, output)
		postorder(root.Right, output)
		*output = append(*output, root.Val)
	}
}
