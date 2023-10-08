package leetcode

// LeetCode：https://leetcode.com/problems/linked-list-cycle/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0141.Linked-List-Cycle/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

// note
// 這題在 TwoPointers 解過，只要利用一塊一慢兩個指標
// 當 slow 跟 fast 撞在一起回傳 true 即可

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
