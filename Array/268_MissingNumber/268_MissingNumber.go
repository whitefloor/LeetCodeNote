package leetcode

// LeetCode：https://leetcode.com/problems/missing-number/description/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0200~0299/0268.Missing-Number/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

// note
// 給定一個陣列，數字範圍是[0,len(array)]，找出陣列中缺少的數字
// follow up : 用 O(1) 的空間跟 O(n) 的時間來解決
// 看解答才想出來的，利用 xor
// 將陣列裡每個數字利用 x^x=0，缺少的另一個 x 利用陣列索引即可
// 注意處理完陣列後，會少做 len(array) 這個數字，所以還要多做一次

func missingNumber(nums []int) int {
	xor, i := 0, 0
	for i = 0; i < len(nums); i++ {
		xor = xor ^ i ^ nums[i]
	}

	return xor ^ i
}
