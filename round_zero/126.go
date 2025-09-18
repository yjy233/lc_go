package roundzero

import (
	"fmt"
	"slices"
)

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordList = append(wordList, beginWord)
	l := len(wordList)
	if l < 2 {
		return [][]string{}
	}
	w2Ind := make(map[string]int)

	for i := 0; i < l; i++ {
		w2Ind[wordList[i]] = i
	}
	fmt.Println(w2Ind)

	_, ok := w2Ind[beginWord]
	if !ok {
		return [][]string{}
	}
	_, ok = w2Ind[endWord]
	if !ok {
		return [][]string{}
	}

	jg := make([][]int, l)
	for i := 0; i < l; i++ {
		jg[i] = make([]int, l)
	}

	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if onlyOne(wordList[i], wordList[j]) {
				jg[i][j] = 1
				jg[j][i] = 1
			}
		}
	}

	for _, i := range jg {
		for _, j := range i {
			fmt.Println(j)
		}
		fmt.Println("")
	}

	res := make([][]string, 0, 100)
	use := make([]bool, l)
	now := w2Ind[beginWord]
	target := w2Ind[endWord]
	path := make([]string, 0, l)

	dfs126(wordList, &res, &use, &path, target, target, now, &jg)

	for i := 0; i < len(res); i++ {
		slices.Reverse(res[i])
	}
	return res
}

func dfs126(wordList []string, res *[][]string, use *[]bool, path *[]string, now, begin, target int, dp *[][]int) {
	fmt.Println(now, target, wordList[now], wordList[target])
	if (*use)[now] {
		return
	}

	if (now != begin) && (*dp)[begin][now] != 0 && (*dp)[begin][now] <= len(*path) {
		return
	}
	(*dp)[begin][now] = len(*path) + 1
	(*dp)[now][begin] = len(*path) + 1

	if now == target {
		if len(*res) == 0 || len(*path) < len((*res)[0]) {
			tmp := make([]string, len(*path), len(*path)+1)
			copy(tmp, *path)
			tmp = append(tmp, wordList[target])

			if len(*res) != 0 && len(*path)+1 == len((*res)[0]) {
				(*res) = append((*res), tmp)
			} else {
				(*res) = (*res)[:0]
				(*res) = append((*res), tmp)
			}
		}

		return
	}

	(*use)[now] = true
	(*path) = append((*path), wordList[now])
	for i := 0; i < len(wordList); i++ {
		if (*use)[i] {
			continue
		}

		dfs126(wordList, res, use, path, i, begin, target, dp)
	}

	(*path) = (*path)[:len(*path)-1]
	(*use)[now] = false

}

func onlyOne(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	l := len(a)
	cnt := 0
	for i := 0; i < l; i++ {
		if a[i] == b[i] {
			continue
		}
		cnt++
		if cnt >= 2 {
			break
		}
	}
	return cnt == 1
}
