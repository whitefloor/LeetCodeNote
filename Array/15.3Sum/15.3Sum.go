package leetcode

import "sort"

// LeetCode：https://leetcode.com/problems/3sum/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0015.3Sum/
// Difficulty：Medium
// Time Complexity：O(n)
// Space Complexity：O(1)

//note
// 從two sum到three sum是複雜不少
// 不過處理這種陣列重要的是最好先排序過會好處理很多
// 而且這題是難在要處理重複解的部分，自己有寫出來計算的部分
// 排序跟解決重複解的地方是上網看的
// 有把參考到的解答稍微改寫一下，宣告的部分就比較簡單了

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result, start, end, length := [][]int{}, 0, 0, len(nums)
	for index := 1; index < length-1; index++ {
		start, end = 0, length-1
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}
		for start < index && end > index {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			addNum := nums[start] + nums[end] + nums[index]
			if addNum == 0 {
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if addNum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return result
}
