package t509

import (
	"fmt"
)

func fib(N int) int {
	if N < 2{
		return N
	}
	cache := make([]int, N+1)
	cache[0], cache[1] = 0, 1
	for i := 2; i <= N; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}
	return cache[N]
}

//var cache map[int]int = nil

//func fib2(N int) int {
//	if N < 2 {
//		return N
//	}
//
//	if cache == nil {
//		cache = make(map[int]int, N+1)
//	}
//
//	if k, ok := cache[N]; ok {
//		return k
//	}
//
//	result := fib2(N-1) + fib2(N-2)
//
//	cache[N] = result
//
//	return result
//}

//func fib3(N int) int {
//
//	var cache map[int]int = make(map[int]int, N+1)
//
//	func
//	reco_fib(N
//	int) int{
//		if N < 2{
//		return N
//	}
//
//		if k, ok := cache[N]; ok{
//		return k
//	}
//
//		result := fib3(N-1) + fib3(N-2)
//
//		cache[N] = result
//	}
//}
//
//	return reco_fib(N)
//}

func ExampleFib() {
	for _, i := range []int{0, 2, 3, 4, 5, 6, 7} {
		fmt.Println(fib(i))
	}

	// Output:
	//0
	//1
	//2
	//3
	//5
	//8
	//13
}
