package main

import (
"fmt"
)

func main() {
	var a, b, c int
	var s string
	// 標準入力からの入力待ち
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d %d", &b, &c)
	fmt.Scanf("%s", &s)

	fmt.Printf("%d %s\n", a+b+c, s)
}
