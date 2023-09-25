package leetcode

// LeetCode：https://leetcode.com/problems/merge-sorted-array/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0088.Merge-Sorted-Array/
// Difficulty：Easy
// Time Complexity：O(m+n)
// Space Complexity：O(1)

// note
// 給你兩個升序排序過後的 int array，nums1、nums2，將其中的 element 升序合併到 nums1中
// 可以認為 nums1 有足夠的空間納入 nums2 的 element
// m＆n 代表 nums1、nums2 的元素數量
// follow up:找出O(m+n)
// 想法是從兩個陣列最後的元素開始比較插入，最後如果 nums2 沒處理完在將 nums2 依序插回 nums1

func merge(nums1 []int, m int, nums2 []int, n int) {
	for p := m + n; m > 0 && n > 0; p-- {
		if nums1[m-1] > nums2[n-1] {
			nums1[p-1] = nums1[m-1]
			m--
		} else {
			nums1[p-1] = nums2[n-1]
			n--
		}
	}
	for ; n > 0; n-- {
		nums1[n-1] = nums2[n-1]
	}
}
