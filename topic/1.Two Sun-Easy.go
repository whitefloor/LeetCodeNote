package topic

// Question

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// Example:
// Given nums = [2, 7, 11, 15], target = 9,
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].

// ---

// 概述

// 最佳解時間複雜度O(n)

// 在int陣列裡找到兩個不重複的元素的元素值相加後等於target
// 不用求多組解，只需要一組解答
// 回傳找到的那兩個元素

func TwoSum(nums []int, target int) []int {
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
