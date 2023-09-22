package leetcode

// LeetCode：https://leetcode.com/problems/remove-duplicates-from-sorted-array/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0026.Remove-Duplicates-from-Sorted-Array/
// Difficulty：
// Time Complexity：
// Space Complexity：

//note

func removeDuplicates(nums []int) int {
	index := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[index] {
			index++
			nums[index] = nums[i]
		}
	}

	return index + 1
}
