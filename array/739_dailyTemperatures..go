package array

// 每日温度

// DailyTemperatures 单调栈实际上就是栈，只是利用了一些巧妙的逻辑，使得每次新元素入栈后，栈内的元素都保持有序（单调递增或单调递减）。
func DailyTemperatures(T []int) []int {
	s := make([]int, 0)
	res := make([]int, len(T))
	for i, v := range T {
		for len(s) > 0 && v > T[s[len(s)-1]] {
			res[s[len(s)-1]] = i - s[len(s)-1]
			s = s[:len(s)-1]
		}
		s = append(s, i)
	}

	return res
}

//DailyTemperatures1 单调栈实际上就是栈，只是利用了一些巧妙的逻辑，使得每次新元素入栈后，栈内的元素都保持有序（单调递增或单调递减）。
func DailyTemperatures1(T []int) []int {
	s := make([]int, 0)
	res := make([]int, len(T))
	for i := len(T) - 1; i >= 0; i-- {
		for len(s) > 0 && T[i] >= T[s[len(s)-1]] {
			s = s[:len(s)-1]
		}

		if len(s) > 0 {
			res[i] = s[len(s)-1] - i
		}

		s = append(s, i)
	}

	return res
}
