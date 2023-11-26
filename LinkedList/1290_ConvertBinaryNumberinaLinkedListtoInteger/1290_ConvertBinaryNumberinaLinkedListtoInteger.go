package leetcode

// LeetCode：https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/description/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/1200~1299/1290.Convert-Binary-Number-in-a-Linked-List-to-Integer/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

// note
// 給定一個單向的鏈結串鏈，其中每個節點儲存的 Val 只有0 or 1，為二進位表示法，求出十進位的值
// 看解答才想出來了，對位元操作不是很熟悉，原本想用遞迴從最後面判斷回來*2
// 不過解答直接從第一位開始往後計算每個節點的值*2也是一樣，而且更簡單不用判斷是不是到鏈表的底了

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	res := 0
	for head != nil {
		res = res*2 + head.Val
		head = head.Next
	}

	return res
}
