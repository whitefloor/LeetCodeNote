package main

// LeetCode:https://leetcode.com/problems/remove-element/description/
// CookBook:https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0027.Remove-Element/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

// note
// 2023/3/21 沒看解答解出來，解題思路跟26題一樣

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	last := 0
	for _, v := range nums {
		if val != v {
			nums[last] = v
			last++
		}
	}

	return last
}
