package main

import (
	"fmt"
)

// ###
// https://atcoder.jp/contests/abs/tasks/practice_1
// ###

func main() {
	var a, b, c int
	var s string

	fmt.Scanf("%d", &a)
	fmt.Scanf("%d %d", &b, &c)
	fmt.Scanf("%s", &s)

	fmt.Printf("%d %s\n", a+b+c, s)
}
