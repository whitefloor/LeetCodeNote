package main

import "log"

// LeetCode:https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
// CookBook:https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0108.Convert-Sorted-Array-to-Binary-Search-Tree/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(n)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	log.Println(nums[len(nums)/2])
	return &TreeNode{
		Val:   nums[len(nums)/2],
		Left:  sortedArrayToBST(nums[:len(nums)/2]),
		Right: sortedArrayToBST(nums[len(nums)/2+1:]),
	}
}

func main() {
	tree := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	log.Println(*tree.Right)
}
