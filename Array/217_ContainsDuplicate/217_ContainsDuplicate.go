package leetcode

// LeetCode：https://leetcode.com/problems/contains-duplicate/description/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0200~0299/0217.Contains-Duplicate/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(n)

// note
// 給一個 int array，如果其中有元素出現過兩次，返回 true，否則 false
// 是挺簡單的，不過一開始有想到使用 XOR 的方法找就不需要用到 map，但是數字並不會兩兩成對出現，就沒辦法

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		if m[nums[i]] > 1 {
			return true
		}
	}

	return false
}
