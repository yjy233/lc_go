package roundzero

import (
	"math/rand"
	"time"
)

func rand13243434() {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	r.Int()

	cnt := make([]int, 5)

	for i := 0; i < 10000; i++ {
		n := r.Intn(5)
		cnt[n]++
	}

	println(cnt)
}
