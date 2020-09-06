package main

import (
	"fmt"
)

// ###
// https://atcoder.jp/contests/abs/tasks/abc086_a
// ###

func main() {
	var a, b int

	fmt.Scanf("%d %d", &a, &b)

	if a*b%2 == 0 {
		// 偶数
		fmt.Print("Even")
	} else {
		// 奇数
		fmt.Print("Odd")
	}
}
