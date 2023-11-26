package leetcode

// LeetCode：https://books.halfrost.com/leetcode/ChapterFour/0700~0799/0704.Binary-Search/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0100.Same-Tree/
// Difficulty：Easy
// Time Complexity：O(logN)
// Space Complexity：O(1)

// note
// 給定一個陣列在其中找到 target，找不到回傳-1
// 經典的二元搜尋
// 要注意的是 for 條件設定跟尋找 middle 的部份即可
// 太久沒寫了有看解答

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		middle := low + (high-low)>>1
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}

	return -1
}
