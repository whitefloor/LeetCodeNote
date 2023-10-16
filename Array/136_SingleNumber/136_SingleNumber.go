package leetcode

// LeetCode：https://leetcode.com/problems/single-number/description/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0136.Single-Number/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

// note
// 給一個整數的陣列，裡面會有成對相同元素，找出有個元素會沒有成對
// 原本想說用排序解，不過似乎無法使用 sort
// 後來看解答用了 XOR 去解，蠻聰明的

func singleNumber(nums []int) int {
	result := 0
	for _, v := range nums {
		result ^= v
	}
	return result
}
