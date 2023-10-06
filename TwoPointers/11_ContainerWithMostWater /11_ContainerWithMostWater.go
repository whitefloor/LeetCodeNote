package leetcode

// LeetCode：https://leetcode.com/problems/container-with-most-water/submissions/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0011.Container-With-Most-Water/
// Difficulty：Medium
// Time Complexity：O(n)
// Space Complexity：O(1)

// note
// 給一個 int array，裡面每個元素都代表牆的高，index 代表地寬，找出在陣列中找出能夠包含最多水是多少
// 還蠻簡單的，只要有兩個點一前一後移動，判斷比較矮的那邊就挪動
// 然後把兩個牆高找出比較矮的那邊，因為比較高的那邊會溢出，之後乘上地寬
// 最後把最高的儲存量做判斷儲存就行

func maxArea(height []int) int {
	i, j, water := 0, len(height)-1, 0

	for i < j {
		h := 0
		weidth := j - i
		if height[i] > height[j] {
			h = height[j]
			j--
		} else {
			h = height[i]
			i++
		}

		w := h * weidth
		if w > water {
			water = w
		}
	}

	return water
}
