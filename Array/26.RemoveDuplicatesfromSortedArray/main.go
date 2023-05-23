package main

import (
	"log"
)

// LeetCode:https://leetcode.com/problems/remove-duplicates-from-sorted-array/
// CookBook:https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0026.Remove-Duplicates-from-Sorted-Array/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

// note
// 原先的思路是利用map統計有多少不同的數字
// 但是題目的測試資料會將 nums 與期待的 array 進行比較，所以需要排列元素
// 另外題目還有要求不要產生空間給新的 array，所以直接在nums進行操作

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last, finder := 0, 0
	for last < len(nums) {
		for nums[last] == nums[finder] {
			finder++
			if finder == len(nums) {
				return last + 1
			}
		}
		nums[last+1] = nums[finder]
		last++
	}

	return last + 1
}

func main() {
	i := removeDuplicates([]int{1, 1, 2})
	log.Println(i)
}
