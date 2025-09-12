package roundzero

func isInterleaveV2(s1 string, s2 string, s3 string) bool {
	res := false
	helper97(s1, s2, s3, 0, 0, 0, &res)
	return res
}

func helper97(s1 string, s2 string, s3 string, a, b, c int, res *bool) {
	if c >= len(s3) {
		*res = true
	}

	if *res {
		return
	}

	if a < len(s1) && s1[a] == s3[c] {
		helper97(s1, s2, s3, a+1, b, c+1, res)
	}

	if *res {
		return
	}

	if b < len(s2) && s2[b] == s3[c] {
		helper97(s1, s2, s3, a, b+1, c+1, res)
	}
}
