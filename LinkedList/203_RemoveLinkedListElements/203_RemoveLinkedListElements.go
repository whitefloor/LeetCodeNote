package leetcode

// LeetCode：https://leetcode.com/problems/remove-linked-list-elements/description/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0200~0299/0203.Remove-Linked-List-Elements/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

// note
// 給定一個 linked list，將符合給定的數值的節點給移除
// 這題其實蠻簡單的，只需要在 head 前面在墊一個假節點指向 head 就行，不過腦子很亂想不出來看了幾個解答才想到
// https://github.com/gaufung/Leetcode-solutions/blob/master/Algorithm/203.%E7%A7%BB%E9%99%A4%E9%93%BE%E8%A1%A8%E5%85%83%E7%B4%A0.go

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	res := &ListNode{}
	res.Next = head
	current := res
	for current.Next != nil {
		if current.Next.Val != val {
			current = current.Next
		} else {
			current.Next = current.Next.Next
		}
	}
	return res.Next
}
