package leetcode

// LeetCode：https://leetcode.com/problems/move-zeroes/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0200~0299/0283.Move-Zeroes/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

// note
// 給你一個 int array 將所有不為0的元素移到 array 最後，同時要保持不為0的元素順序
// 蠻簡單的，自己直接寫出來，有看了一下別人寫的優化 code

func moveZeroes(nums []int) {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
}
