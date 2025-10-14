package roundzero

import (
	"sort"
)

/*

急的ind
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	//fmt.Println(l)

	if l%2 == 1 {
		return float64(helper4(nums1, nums2, l/2))
	}

	f0 := float64(helper4(nums1, nums2, l/2-1))
	f1 := float64(helper4(nums1, nums2, l/2))
	//fmt.Println(f0, f1)
	return (f0 + f1) / 2
}

func helper4(nums1, nums2 []int, ind int) int {
	l0 := len(nums1)
	l1 := len(nums2)

	if l0 == 0 {
		return nums2[ind]
	}
	if l1 == 0 {
		return nums1[ind]
	}

	if ind <= 0 {
		return min(nums1[0], nums2[0])
	}

	if ind == 1 {
		tmp := make([]int, 0, 4)
		tmp = append(tmp, nums1[0])
		tmp = append(tmp, nums2[0])
		if l0 > 1 {
			tmp = append(tmp, nums1[1])
		}
		if l1 > 1 {
			tmp = append(tmp, nums2[1])
		}

		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})
		return tmp[1]
	}

	newInd := max(min(l1/2-1, l0/2-1, ind/2), 0)
	if nums1[newInd] < nums2[newInd] {
		newInd = max(newInd, 1)
		return helper4(nums1[newInd:], nums2, ind-newInd)
	}

	newInd = max(newInd, 1)
	return helper4(nums1, nums2[newInd:], ind-newInd)
}
