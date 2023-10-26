package leetcode

// LeetCode：https://leetcode.com/problems/maximum-average-subarray-i/description/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0600~0699/0643.Maximum-Average-Subarray-I/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：

// note
// 給定一個 int array 在其中找到四個連續元素的子陣列，返回最大的平均值
// 原本是想要把透過計算找到每個連續的子陣列計算
// 但是題目要求連續的元素，所以計算出[0,k]的值後，往後滑動只需要 sum = sum - nums[i-k] + nums[k] 即可找出最大
// 看解答想出來的

func findMaxAverage(nums []int, k int) float64 {
	max := 0
	for i := 0; i < k; i++ {
		max += nums[i]
	}

	sum := max
	for i := k; i < len(nums); i++ {
		sum = sum - nums[i-k] + nums[i]
		if sum > max {
			max = sum
		}
	}

	return float64(max) / float64(k)
}
