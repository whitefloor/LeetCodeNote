package leetcode

// LeetCode：https://leetcode.com/problems/majority-element/description/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0169.Majority-Element/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

// note
// 給一個 int 陣列，從其中找出出現次數最多次的元素，而該元素出現次數必定會超過陣列長度的一半
// 其實是蠻簡單的，用 map 就可以解，但是 follow up 要求 space complexity O(1)
// 後來去看解答才知道空間複雜度 O(1) 的寫法

func majorityElement(nums []int) int {
	result, count := nums[0], 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			result, count = nums[i], 1
		} else {
			if result == nums[i] {
				count++
			} else {
				count--
			}
		}
	}

	return result
}
