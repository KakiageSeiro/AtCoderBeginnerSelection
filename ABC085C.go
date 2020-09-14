package main

import (
	"fmt"
)

// ###
// https://atcoder.jp/contests/abs/tasks/abc085_c
// ###

func main() {

	var n, y int
	fmt.Scanf("%d %d", &n, &y)

	for i := 0; i <= n; i++ { // 1万円札
		for j := 0; j <= n-i; j++ { // 5千円札(1万円札が一枚以上選択している場合は、その分ループしなくていいので条件式がn-i)
			k := n - i - j // 千円札
			if y == i*10000+j*5000+k*1000 {
				fmt.Printf("%d %d %d", i, j, k)
				return
			}
		}
	}

	// 見つからなかった場合
	fmt.Print("-1 -1 -1")
}
