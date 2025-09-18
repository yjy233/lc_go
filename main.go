package main

import (
	"fmt"

	roundzero "github.com/yjy233/lc_go/round_zero"
)

func main() {
	nums := []int{
		1, 2, 3, 4, 5, 6, 7,
	}

	roundzero.Rotate189(nums, 3)
	fmt.Println(nums)
}
