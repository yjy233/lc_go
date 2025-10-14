package roundzero

import "fmt"

func getHint(secret string, guess string) string {
	n := len(guess)

	c2Cnt := make(map[byte]int)
	use := make([]bool, n)

	bulls := 0
	cows := 0
	for i := range secret {
		if i < n && secret[i] == guess[i] {
			bulls++
			use[i] = true
			continue
		}

		c2Cnt[secret[i]] += 1
	}

	for j := range guess {
		if use[j] {
			continue
		}

		if c2Cnt[guess[j]] > 0 {
			cows += 1
			c2Cnt[guess[j]] -= 1
		}
	}

	return fmt.Sprintf("%dA%dB", bulls, cows)
}
