package leetcode

// LeetCode：https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0121.Best-Time-to-Buy-and-Sell-Stock/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：

// note
// 給定一個整數陣列，裡面是買賣股票的價格，在其中找出利潤最高是多少
// 注意：不能第二天買，第一天賣
// 看解答寫出來的，自己寫的測資沒全過

func maxProfit(prices []int) int {
	min, maxProfit := prices[0], 0
	for i := 0; i < len(prices); i++ {
		if prices[i]-min > maxProfit {
			maxProfit = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}

	return maxProfit
}
