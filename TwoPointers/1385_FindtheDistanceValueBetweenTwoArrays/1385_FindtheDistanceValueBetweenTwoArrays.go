package leetcode

// LeetCode：https://leetcode.com/problems/find-the-distance-value-between-two-arrays/description/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/1300~1399/1385.Find-the-Distance-Value-Between-Two-Arrays/
// Difficulty：Easy
// Time Complexity：
// Space Complexity：O(1)

// note
// 敘述超爛
// 意思是給你兩個陣列，如果 arr1 中有任何元素與 arr2 的絕對值小於 d，則該元素不計數

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	res := 0
	for i := range arr1 {
		for j := range arr2 {
			if abs(arr1[i]-arr2[j]) <= d {
				break
			}
			if j == len(arr2)-1 {
				res++
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
