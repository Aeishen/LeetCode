package array

// 下一个更大元素 II

//NextGreaterElements 单调栈实际上就是栈，只是利用了一些巧妙的逻辑，使得每次新元素入栈后，栈内的元素都保持有序（单调递增或单调递减）。
// 一般是通过 % 运算符求模（余数），来获得环形特效：
func NextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	s := make([]int, 0)
	n := len(nums)
	for i := 2*n - 1; i >= 0 && n > 0; i-- {
		for len(s) > 0 && nums[i%n] >= s[len(s)-1] {
			s = s[:len(s)-1]
		}

		if len(s) > 0 {
			res[i%n] = s[len(s)-1]
		} else {
			res[i%n] = -1
		}

		s = append(s, nums[i%n])
	}

	return res
}
