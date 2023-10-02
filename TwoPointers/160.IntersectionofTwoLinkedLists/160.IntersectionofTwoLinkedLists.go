package leetcode

// LeetCode：https://leetcode.com/problems/intersection-of-two-linked-lists/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0160.Intersection-of-Two-Linked-Lists/
// Difficulty：Easy
// Time Complexity：
// Space Complexity：

// note
// 給定 A/B 兩個 linked list 從中找出兩個指向同一個 node 的交叉點，交叉點值不為0
// follow up 使用 t:O(m+n) s:(1) 的方法解決此問題
// 後來是看解答，他將兩個不一樣長度的 linked list 拼成一樣長（a+b），然後對兩組 list 進行比對

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a := headA
	b := headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}
