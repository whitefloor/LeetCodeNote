package Template

// Go 1.21 版本後加入 max()、min() 內建方法，在 LeetCode 上可用
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
