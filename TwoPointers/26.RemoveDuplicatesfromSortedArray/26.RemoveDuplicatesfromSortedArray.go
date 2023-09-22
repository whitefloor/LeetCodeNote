package leetcode

// LeetCode：https://leetcode.com/problems/remove-duplicates-from-sorted-array/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0026.Remove-Duplicates-from-Sorted-Array/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

// note
// 題目大意是給一個排序過後的 int array，並將重複的去除並返回沒有重複的長度，然後不用宣告一個新的陣列
// 不重複的值要從 array[0] 開始排序
// 重複的值不用真的刪掉
// 只要給一個從0開始的 index，後面遇到不重複的 index 就+1，把當前不重複的值賦與到 index
// 最後只要記得把 index+1 輸出就可以

func removeDuplicates(nums []int) int {
	index := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[index] {
			index++
			nums[index] = nums[i]
		}
	}

	return index + 1
}
