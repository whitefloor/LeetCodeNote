package leetcode

// LeetCode：https://leetcode.com/problems/trapping-rain-water/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0042.Trapping-Rain-Water/
// Difficulty：Hard
// Time Complexity：O(n)
// Space Complexity：O(1)

// note
// 這題跟11題有點類似，不過不是直接算出左右高牆內的水，而是要計算左右高牆內扣除中間的牆所能包含的水量
// 看解答才知道為什麼，而且解法很厲害，自己有想到扣除障礙物的部份，不過解法還另外記錄了左右最高牆的高度去做計算
// 所以最後只需要 time O(n) 就能計算完成
// 解法一樣是從左右使用雙指針位移，比較矮的那邊就要進行移動，同時記錄左右最高牆是多少
// 當遇到比較矮的牆就將最高牆減去比較矮的牆就是該段的儲水量，這樣 每個元素的位置就會分段計算
// 遇到比較高的牆時就更新最高牆的紀錄

func trap(height []int) int {
	res, left, right, maxLeft, maxRight := 0, 0, len(height)-1, 0, 0
	for left < right {
		if height[left] < height[right] {
			if height[left] > maxLeft {
				maxLeft = height[left]
			} else {
				res += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] > maxRight {
				maxRight = height[right]
			} else {
				res += maxRight - height[right]
			}
			right--
		}
	}

	return res
}
