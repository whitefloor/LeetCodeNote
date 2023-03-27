package main

// LeetCode:https://leetcode.com/problems/merge-sorted-array/
// CookBook:https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0088.Merge-Sorted-Array/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

// note
// nums1 長度為兩個陣列中元素的總和
// 所以從最陣列最後一位開始比較兩個陣列的值在插入
// 最後要防範最差的情況，都是nums1 > nums2
// 所以最後還要特別處理 nums2 插入到陣列內

func merge(nums1 []int, m int, nums2 []int, n int) {
	for p := m + n; m > 0 && n > 0; p-- {
		if nums1[m-1] <= nums2[n-1] {
			nums1[p-1] = nums2[n-1]
			n--
		} else {
			nums1[p-1] = nums1[m-1]
			m--
		}
	}
	for ; n > 0; n-- {
		nums1[n-1] = nums2[n-1]
	}
}
