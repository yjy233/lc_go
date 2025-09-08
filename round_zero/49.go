package roundzero

import "sort"

func groupAnagrams(strs []string) [][]string {
	key2StrList := make(map[string][]string)

	for _, term := range strs {
		key := genKeyForGroupAnagrams(term)
		key2StrList[key] = append(key2StrList[key], term)
	}

	realRes := make([][]string, 0, len(key2StrList))

	for _, v := range key2StrList {
		realRes = append(realRes, v)
	}
	return realRes
}

func genKeyForGroupAnagrams(str string) string {
	byteL := ([]byte)(str)

	sort.Slice(byteL, func(i, j int) bool {
		return byteL[i] < byteL[j]
	})

	return string(byteL)
}
