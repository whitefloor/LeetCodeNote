package leetcode

// LeetCode：https://leetcode.com/problems/linked-list-cycle/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0141.Linked-List-Cycle/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

// note
// 有參考解答，對 linked list 還不太熟悉不知道該怎麼解
// 這題主要是要在 linked list 中尋找有沒有具有 ring 結構的迴圈，follow up 是使用空間複雜度 O(1) 來解決
// 一開始是想使用 hash map，但在 leetcode 跟 cookbook 上看到 Floyd's Cycle-Finding Algorithm
// 原理就是利用兩個 pointer 在跑，一個跑比較快，另一個跑比較慢，跑到最後一定會碰到，空間複雜度為 O(1)

// ListNode Definition for singly-linked list.
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
