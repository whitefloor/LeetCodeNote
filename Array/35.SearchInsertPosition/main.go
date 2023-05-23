package main

// LeetCode:https://leetcode.com/problems/search-insert-position/description/
// CookBook:https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0035.Search-Insert-Position/
// Difficulty：Easy
// Time Complexity：O(logN)
// Space Complexity：O(1)

func searchInsert(nums []int, target int) int {
	mid, left, right := 0, 0, len(nums)-1

	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] >= target) {
				return mid + 1
			}
			left = mid + 1
		}
	}

	return mid
}

func main() {
}
