package roundzero

func findMedianSortedArraysV1(nums1 []int, nums2 []int) float64 {

	l := len(nums1) + len(nums2)
	if l%2 == 1 {
		return float64(findN(l/2, nums1, nums2))
	}
	r := float64(findN(l/2-1, nums1, nums2))
	rr := float64(findN(l/2, nums1, nums2))
	return (rr + r) / 2.0
}

func findN(ind int, nums1, nums2 []int) int {

	if len(nums1) == 0 {
		return nums2[ind]
	}

	if len(nums2) == 0 {
		return nums1[ind]
	}

	if nums1[0] > nums2[0] {
		nums1, nums2 = nums2, nums1
	}

	if ind == 0 {
		return nums1[0]
	}

	if nums1[len(nums1)-1] <= nums2[0] {
		if ind < len(nums1) {
			return nums1[ind]
		}
		return nums2[ind-len(nums1)]
	}

	rml := min(len(nums1)/2-1, len(nums2)/2-1, ind/2, len(nums1)-1, len(nums2)-1)
	rml = max(0, rml)

	if nums1[rml] < nums2[rml] {
		if rml == 0 {
			rml++
		}
		return findN(ind-rml, nums1[rml:], nums2)
	}

	if rml == 0 {
		rml++
	}
	return findN(ind-rml, nums1, nums2[rml:])
}
