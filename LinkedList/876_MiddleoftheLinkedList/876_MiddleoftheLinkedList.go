package leetcode

// LeetCode：https://leetcode.com/problems/middle-of-the-linked-list/description/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0800~0899/0876.Middle-of-the-Linked-List/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

// note
// 給你一個 linked list，找出在其中間的節點，如果中間正巧有兩個節點，則返回後面的那個
// 蠻簡單的，Floyd演算法，一快一慢就可以解出來，沒看解答

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
