package leetcode

// LeetCode：https://leetcode.com/problems/flipping-an-image/description/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0800~0899/0832.Flipping-an-Image/?
// Difficulty：Easy
// Time Complexity：
// Space Complexity：O(1)

// note
// 給你個二進制矩陣，要求將每個矩陣中的元素前後反轉後在進行not操作
// 蠻簡單的，其實只是要 for each 翻轉而已，沒看解答就寫出來了，考驗基本陣列操作
// 不過我的 not 操作比解答還要好

func flipAndInvertImage(image [][]int) [][]int {
	for r := 0; r < len(image); r++ {
		i, j := 0, len(image[r])-1
		for i < j {
			image[r][i], image[r][j] = image[r][j], image[r][i]
			i++
			j--
		}
		for k := range image[r] {
			image[r][k] ^= 1
		}
	}

	return image
}
