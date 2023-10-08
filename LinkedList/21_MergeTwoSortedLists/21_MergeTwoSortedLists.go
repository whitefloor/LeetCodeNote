package leetcode

// LeetCode：https://leetcode.com/problems/merge-two-sorted-lists/
// CookBook：解題的時候網站爆了，待補
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

// note
// 給兩個排序過的 linked list，將這兩串合併後升序輸出
// 原本比較偏向解法1的想法，後來看到了用Recursion（遞迴）的解法
// 看解答後寫出來的

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 == nil {
		current.Next = list2
	} else {
		current.Next = list1
	}

	return result.Next
}

func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists2(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists2(list1, list2.Next)

	return list2
}
