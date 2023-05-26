package leetcode

// LeetCode：https://leetcode.com/problems/binary-tree-inorder-traversal/submissions/957513623/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0094.Binary-Tree-Inorder-Traversal/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：

//note
// 標準的DFS，不過是中序遍歷，順序是左、根、右
// https://web.ntnu.edu.tw/~algo/BinaryTree.html
// 參考解答學DFS的，以前沒處理過

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

func inorderTraversal(root *TreeNode) []int {
	var result []int
	inorder(root, &result)
	return result
}

func inorder(root *TreeNode, output *[]int) {
	if root != nil {
		inorder(root.Left, output)
		*output = append(*output, root.Val)
		inorder(root.Right, output)
	}
}
