package leetcode

// LeetCode：https://leetcode.com/problems/remove-element/submissions/
// CookBook：https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0026.Remove-Duplicates-from-Sorted-Array/
// Difficulty：Easy
// Time Complexity：O(n)
// Space Complexity：O(1)

// note
// 這題跟26題基本上一樣，只要把 array 中與 val 相同的去除，返回不等於 val 的 element 個數即可
// 不等於 val 的值要從 array[0] 開始排序
// 要求刪除 element 的不用真的刪除
// 只要把不等於 val 的值從 array[0] 開始排序，最後輸出 index 即可

func removeElement(nums []int, val int) int {
	index := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}

	return index
}
