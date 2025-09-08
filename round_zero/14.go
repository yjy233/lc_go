package roundzero

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	res := strs[0]
	for i := 1; i < len(strs); i++ {
		s := strs[i]
		tmp := []byte{}
		for j := 0; j < min(len(res), len(s)); j++ {
			if s[j] != res[j] {
				break
			}

			tmp = append(tmp, s[j])
		}

		res = string(tmp)
		if res == "" {
			break
		}
	}
	return res
}
