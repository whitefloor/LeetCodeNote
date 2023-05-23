package leetcode

// Question

// Given an integer x, return true if x is palindrome integer.
// An integer is a palindrome when it reads the same backward as forward.
// For example, 121 is a palindrome while 123 is not.

// Example 1:

// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.

// Example 2:

// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

// Example 3:

// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

// Follow up: Could you solve it without converting the integer to a string?

// 概述

// 判斷一個數字是否為回文(palindrome)從左到右，從右到左的順序都相同

// Follow up：能夠在不轉換成字串的情況下完成嗎？

// 筆記：從cookbook看來的解法，有優化減少一些if判斷跟不必要的cap宣告，較為精簡

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	arr := []int{}
	for x > 0 {
		arr = append(arr, x%10)
		x = x / 10
	}

	rightSide := len(arr) - 1
	for i, j := 0, rightSide; i <= j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}

	return true
}
