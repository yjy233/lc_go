package roundzero

func findMedianSortedArraysV3(nums1 []int, nums2 []int) float64 {

	l1 := len(nums1)
	l2 := len(nums2)

	if (l1+l2)%2 == 1 {
		// 1
		// 2 3
		return float64(find4(nums1, nums2, (l1+l2)/2))
	}

	f1 := float64(find4(nums1, nums1, (l1+l2)/2-1))
	f2 := float64(find4(nums1, nums1, (l1+l2)/2))
	return (f1 + f2) / 2
}

func find4(nums1, nums2 []int, idx int) int {
	l1 := len(nums1)
	l2 := len(nums2)

	if l1 == 0 {
		return nums2[idx]
	}

	if l2 == 0 {
		return nums1[idx]
	}

	if idx == 0 {
		return min(nums1[0], nums2[0])
	}

	midIdx := min(l1/2-1, l2/2-1)
	midIdx = min(idx/2, midIdx)
	midIdx = max(0, midIdx)

	if nums1[midIdx] < nums2[midIdx] {
		if midIdx == 0 {
			midIdx = 1
		}

		return find4(nums1[midIdx:], nums2, idx-midIdx)
	}

	if midIdx == 0 {
		midIdx = 1
	}

	return find4(nums1, nums2[midIdx:], idx-midIdx)
}
