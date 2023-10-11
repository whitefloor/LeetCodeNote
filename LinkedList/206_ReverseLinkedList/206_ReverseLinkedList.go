package leetcode

// LeetCode：https://leetcode.com/problems/reverse-linked-list/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0200~0299/0206.Reverse-Linked-List/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

// note
// 給定一個單向鏈結串鏈，反轉此串鏈即可
// 反轉要多練習，有看了解答，而且要注意 var 跟 := 初始化的差别

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var res *ListNode
	for head != nil {
		next := head.Next
		head.Next = res
		res = head
		head = next
	}

	return res
}
