package leetcode

// LeetCode：https://leetcode.com/problems/intersection-of-two-arrays/description/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0300~0399/0349.Intersection-of-Two-Arrays/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(n)

// note
// 給你兩個陣列找除其中重複的元素，每個重複的元素只在回傳陣列中輸出一次，不在意輸出順序
// 一開始想用 hash map 解，想看看有沒有更好的解法，不過好像大家都這樣，那就很簡單
// 後來有看別人寫的改進 code，for 迴圈可以少用一次，也更簡單

func intersection(nums1 []int, nums2 []int) []int {
	res := []int{}
	m := map[int]bool{}

	for _, v := range nums1 {
		m[v] = true
	}
	for _, v := range nums2 {
		if m[v] {
			delete(m, v)
			res = append(res, v)
		}
	}

	return res
}
