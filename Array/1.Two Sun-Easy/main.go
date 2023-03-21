package main

// LeetCode:https://leetcode.com/problems/two-sum/description/
// CookBook:https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0001.Two-Sum/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(n)

// TwoSum1 O(n^2)
func TwoSum1(nums []int, target int) []int {
	var answer []int

	for i := 0; i <= len(nums)-2; i++ {
		for j := i + 1; j <= len(nums)-1; j++ {
			if nums[i]+nums[j] == target {
				answer = append(answer, i, j)
				break
			}
		}
	}

	return answer
}

// TwoSum2 O(n)
func TwoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if _, ok := m[target-v]; ok {
			return []int{m[target-v], i}
		}
		m[v] = i
	}

	return nil
}
