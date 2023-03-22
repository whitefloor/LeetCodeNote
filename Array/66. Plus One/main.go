package main

// LeetCode:https://leetcode.com/problems/plus-one/
// CookBook:https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0066.Plus-One/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] != 10 {
			return digits
		}
		digits[i] = 0
	}
	digits[0] = 1
	digits = append(digits, 0)
	return digits
}
