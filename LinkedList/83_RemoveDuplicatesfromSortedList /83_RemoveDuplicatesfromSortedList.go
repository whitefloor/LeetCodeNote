package leetcode

// LeetCode：https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0083.Remove-Duplicates-from-Sorted-List/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

// note
// 給排序過後的 linked list，去掉重複的數值並返回
// 判斷去重複的思路很簡單，不過對資料結構不熟悉，條件設錯了，看解答才想出來為什麼
// 而且解答返回的時候記得要記錄標頭，會比較好處理
// 設置一個標頭紀錄一開始的位置，接著開始往後移動，只要遇到是一樣的就把更後面的節點值抓來當下一個節點
// 如果不一樣的話就只要移動當前位置即可

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates1(head *ListNode) *ListNode {
	current := head
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	} else if head.Val == head.Next.Val {
		head = deleteDuplicates2(head.Next)
		return head
	} else {
		head.Next = deleteDuplicates2(head.Next)
		return head
	}
}
