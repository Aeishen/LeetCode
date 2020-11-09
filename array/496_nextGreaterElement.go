package array

//方法一：单调栈
//我们可以忽略数组 nums1，先对将 nums2 中的每一个元素，求出其下一个更大的元素。
//随后对于将这些答案放入哈希映射（HashMap）中，再遍历数组 nums1，并直接找出答案。
//对于 nums2，我们可以使用单调栈来解决这个问题。 我们首先把第一个元素 nums2[1] 放入栈，
//随后对于第二个元素 nums2[2]，如果 nums2[2] > nums2[1]，那么我们就找到了 nums2[1] 的下一个更大元素 nums2[2]，
//此时就可以把 nums2[1] 出栈并把 nums2[2] 入栈；如果 nums2[2] <= nums2[1]，我们就仅把 nums2[2] 入栈。
//对于第三个元素 nums2[3]，此时栈中有若干个元素，那么所有比 nums2[3] 小的元素都找到了下一个更大元素（即 nums2[3]），
//因此可以出栈，在这之后，我们将 nums2[3] 入栈，以此类推。
//可以发现，我们维护了一个单调栈，栈中的元素从栈顶到栈底是单调不降的。
//当我们遇到一个新的元素 nums2[i] 时，我们判断栈顶元素是否小于 nums2[i]，如果是，那么栈顶元素的下一个更大元素即为 nums2[i]，
//我们将栈顶元素出栈。重复这一操作，直到栈为空或者栈顶元素大于 nums2[i]。此时我们将 nums2[i] 入栈，保持栈的单调性，
//并对接下来的 nums2[i + 1], nums2[i + 2] ... 执行同样的操作。

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	s := make([]int, 0)
    for _, v := range nums2{
    	for len(s) > 0 && v > s[len(s) - 1]{
			m[s[len(s) - 1]] = v
			s = s[:len(s) - 1]
		}

		s = append(s,v)
	}
	for i, v := range nums1{
		if v, ok := m[v]; ok {
			nums1[i] = v
		}else{
			nums1[i] = -1
		}

	}
	return nums1
}

func nextGreaterElement1(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	s := make([]int, 0)
	for i := len(nums2) - 1; i >= 0; i--{
		for len(s) > 0 && s[len(s) - 1] <= nums2[i]{
			s = s[:len(s) - 1]
		}
		if len(s) == 0 {
			m[nums2[i]] = -1
		}else{
			m[nums2[i]] = s[len(s) - 1]
		}
		s = append(s,nums2[i])
	}

	for i, v := range nums1{
		nums1[i] = m[v]
	}

	return nums1
}
