package roundzero

func sortColors(nums []int) {
	l := len(nums)

	zeroB := 0
	for i := 0; i < l; i++ {
		if nums[zeroB] != 0 {
			break
		}
		zeroB++
	}

	twoB := l - 1
	for i := l - 1; i >= 0; i-- {
		if nums[twoB] != 2 {
			break
		}

		twoB--
	}

	if zeroB >= twoB {
		return
	}

	i := zeroB
	for i < l && i >= zeroB && i <= twoB {
		if i < zeroB {
			i++
			continue
		}

		if nums[i] == 1 {
			i++
			continue
		}

		if nums[i] == 2 {
			nums[i], nums[twoB] = nums[twoB], nums[i]
			twoB--
			continue
		}

		if nums[i] == 0 {
			if i == zeroB {
				i++
				continue
			}
			if nums[zeroB] == 0 {
				zeroB++
				continue
			}

			nums[zeroB], nums[i] = nums[i], nums[zeroB]
			zeroB++
		}
	}
}
