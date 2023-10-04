package leetcode

import "log"

// LeetCode：https://leetcode.com/problems/palindrome-linked-list/description/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0200~0299/0234.Palindrome-Linked-List/
// Difficulty：Easy
// Time Complexity：
// Space Complexity：

// note
// 在 linked list 中判斷是不是回文（Palindrome）
// follow up 用 O(n) time and O(1) space
// 一開始有想到用陣列的方法解決，純用 linked list 的要看解答而且程式碼會複雜很多
// 需要先找出陣列切半的位置，接著反轉陣列，最後進行比較

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome1(head *ListNode) bool {
	arr := []int{}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	i, j := 0, len(arr)-1
	for i < j {
		if arr[i] != arr[j] {
			return false
		}
		i++
		j--
	}

	return true
}

// ======

func isPalindrome2(head *ListNode) bool {
	sublist := findMiddle(head)
	sublist = reverseList(sublist)
	return compareList(head, sublist)
}

func findMiddle(head *ListNode) *ListNode {
	fast, slow := head, head
	// 用這種方法正好比較快的抵達最後，比較慢的正巧到一半
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		// 暫存下一個節點
		nextNode := curr.Next
		if nextNode != nil {
			log.Println("nextNode:", nextNode)
		} else {
			log.Println("nextNode nil")
		}
		//
		curr.Next = prev
		if curr.Next != nil {
			log.Println("curr.Next:", curr.Next)
		} else {
			log.Println("curr.Next nil")
		}
		//
		prev = curr
		if prev != nil {
			log.Println("prev:", prev)
		} else {
			log.Println("prev nil")
		}
		// 指定下一個節點
		curr = nextNode
		if curr != nil {
			log.Println("curr:", curr)
		} else {
			log.Println("curr nil")
		}
	}
	log.Println(prev)

	return prev
}

func compareList(a, b *ListNode) bool {
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}

	return true
}
