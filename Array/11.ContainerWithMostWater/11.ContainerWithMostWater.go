package leetcode

// LeetCode：https://leetcode.com/problems/container-with-most-water/description/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0011.Container-With-Most-Water/
// Difficulty：Medium
// Time Complexity：O(n)
// Space Complexity：O(1)

//note
// 這題很久以前有練習解過一次
// 不過這次在2023/5/22就直接寫出來了，沒看任何參考
// 思路是從陣列的頭跟尾開始判斷誰比較高
// 能儲存的水就高*寬就好
// 判斷完了就把比較低的那邊不用繼續判斷，把index往前或往後挪就好

func maxArea(height []int) int {
	start, end, mostWater, newWater := 0, len(height)-1, 0, 0
	for start < end {
		if height[start] >= height[end] {
			newWater = height[end] * (end - start)
			end--
		} else {
			newWater = height[start] * (end - start)
			start++
		}

		if newWater > mostWater {
			mostWater = newWater
		}
	}

	return mostWater
}
