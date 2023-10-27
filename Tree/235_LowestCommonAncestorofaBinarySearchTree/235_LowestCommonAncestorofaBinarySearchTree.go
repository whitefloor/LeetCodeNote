package leetcode

// LeetCode：https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0200~0299/0235.Lowest-Common-Ancestor-of-a-Binary-Search-Tree/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：

// note
// 給定一個 BST（二元搜索樹）即 p、q 找出最近的共同祖先節點（lowest common ancestor (LCA)）
// 這提要利用二元搜索樹的特性去找，左節點比當前節點小，右節點比當前節點大
// 看解答才寫出來的，這提要了解 BST、LCA

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return root
}
