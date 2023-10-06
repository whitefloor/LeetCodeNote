package leetcode

import "sort"

// LeetCode：https://leetcode.com/problems/3sum/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0015.3Sum/
// Difficulty：Medium
// Time Complexity：O(n^2)
// Space Complexity：O(n)

// note
// 給一個 integer array，從中找出三個元素和為0的組合
// 組合元素不能重複，例如[1,1,2],[2,1,1]，但排序不影響輸出
// 這題主要的問題是要去重複的組合
// 解法是固定一個點從一直挪到陣列的最後一個元素，在從兩側不斷的去做相加還有判斷去重複，還蠻難的

func threeSum(nums []int) [][]int {
	res, lenth := [][]int{}, len(nums)
	sort.Ints(nums)
	for index := 1; index < lenth-1; index++ {
		start, end := 0, lenth-1
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}
		for start < index && index < end {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < lenth-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			addNum := nums[start] + nums[index] + nums[end]
			if addNum == 0 {
				res = append(res, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if addNum > 0 {
				end--
			} else {
				start++
			}
		}
	}

	return res
}
