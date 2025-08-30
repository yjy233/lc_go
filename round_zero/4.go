package roundzero

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)

	if l%2 == 1 {
		return helper(nums1, nums2, l/2)
	}

	f1 := helper(nums1, nums2, l/2)
	f2 := helper(nums1, nums2, l/2-1)

	return (f1 + f2) / 2.0
}

func helper(nums1 []int, nums2 []int, ind int) float64 {
	//fmt.Println(nums1, nums2, ind)
	if len(nums1) == 0 {
		return float64(nums2[ind])
	}

	if len(nums2) == 0 {
		return float64(nums1[ind])
	}

	if ind == 0 {
		return float64(min(nums1[0], nums2[0]))
	}

	if ind == 1 {
		l := []int{nums1[0], nums2[0]}
		if len(nums1) > 1 {
			l = append(l, nums1[1])
		}
		if len(nums2) > 1 {
			l = append(l, nums2[1])
		}

		sort.Slice(l, func(i, j int) bool {
			return l[i] < l[j]
		})
		return float64(l[1])
	}

	span := min(ind/2, len(nums1)/2)
	span = min(span, len(nums2)/2)

	if span == 0 {
		if nums1[0] < nums2[0] {
			nums1 = nums1[1:]
		} else {
			nums2 = nums2[1:]
		}

		return helper(nums1, nums2, ind-1)
	}

	if nums1[span] < nums2[span] {
		return helper(nums1[span:], nums2, ind-span)
	}

	return helper(nums1, nums2[span:], ind-span)
}
