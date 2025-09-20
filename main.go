package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
给定一个整数数组 nums 和一个整数 k ，返回其中元素之和可被 k 整除的非空 子数组 的数目。
nums = [4,5,0,-2,-3,1], k = 5
输出:7
有 7 个子数组满足其元素之和可被 k = 5 整除：
[4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]
*/
func main() {
	i := int32(0)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for atomic.LoadInt32(&i) <= 10 {
			t := atomic.LoadInt32(&i)
			if t%2 == 1 {
				fmt.Println("1 ", t)
				atomic.AddInt32(&i, 1)
			}

		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for atomic.LoadInt32(&i) <= 10 {
			t := atomic.LoadInt32(&i)
			if t%2 == 0 {
				fmt.Println("2 ", t)
				atomic.AddInt32(&i, 1)
			}
		}
	}()

	wg.Wait()

	for 1 > 0 {

	}
}
