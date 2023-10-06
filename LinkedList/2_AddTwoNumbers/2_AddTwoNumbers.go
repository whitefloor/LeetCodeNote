package main

// LeetCode：https://leetcode.com/problems/add-two-numbers/submissions/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0002.Add-Two-Numbers/
// Difficulty：Medium
// Time Complexity：O(n)
// Space Complexity：O(1)

// note
// 給你兩個 linked list 裡面每個節點裝的都是一位數字，然後事正常數字的倒裝，例如123 : 3>2>1
// 最後要把兩個連結串列相加後返回同樣結構輸出
// 解法蠻厲害的，做逆向加法，還多指定一個虛擬的 head，最後輸出 head.Next 就是正確答案

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	n1, n2, carry, current := 0, 0, 0, head

	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		} else {
			n1 = 0
		}

		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		} else {
			n2 = 0
		}

		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}

	return head.Next
}
