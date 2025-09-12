package roundzero

func merge88(nums1 []int, m int, nums2 []int, n int) {
	for i := m - 1; i >= 0; i-- {
		nums1[i+n] = nums1[i]
	}

	i := 0
	i0 := n
	i1 := 0

	for i < m+n {
		if i0 >= m+n {
			nums1[i] = nums2[i1]
			i1++
			i++
			continue
		}

		if i1 >= n {
			nums1[i] = nums1[i0]
			i0++
			i++
			continue
		}

		if nums1[i0] < nums2[i1] {
			nums1[i] = nums1[i0]
			i0++
			i++
			continue
		}

		nums1[i] = nums2[i1]
		i1++
		i++
	}
}
