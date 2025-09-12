package roundzero

import "sort"

func subsetsWithDup(nums []int) [][]int {
	l := len(nums)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	res := make([][]int, 0, 1<<l)
	tmp := make([][]int, 0, 1<<l)
	res = append(res, []int{})

	for i := 0; i < l; i++ {
		cnt := 0
		for j := i - 1; j >= 0; j-- {
			if nums[j] != nums[i] {
				break
			}

			cnt++
		}

		for _, terms := range res {

			tmp = append(tmp, terms)

			need := true
			if cnt != 0 {
				tCnt := 0
				for j := len(terms) - 1; j >= 0; j-- {
					if terms[j] != nums[i] {
						break
					}

					tCnt++
				}

				need = (tCnt >= cnt)
			}

			if need {
				newTerms := make([]int, len(terms), len(terms)+1)
				copy(newTerms, terms)
				newTerms = append(newTerms, nums[i])
				tmp = append(tmp, newTerms)
			}

		}

		tmp, res = res, tmp
		tmp = tmp[:0]
	}
	return res
}
