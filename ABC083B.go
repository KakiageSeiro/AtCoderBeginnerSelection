package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ###
// https://atcoder.jp/contests/abs/tasks/abc083_b
// ###

func main() {

	var n, a, b  int
	fmt.Scanf("%d %d %d", &n, &a, &b)

	resultSum := 0
	for i := 1; i <= n; i++ {
		// 桁ごとの数値を合計する
		var sum int = 0
		itoa := strconv.Itoa(i)
		slice := strings.Split(itoa, "")
		for i2 := 0; i2 < len(slice); i2++ {
			atoi, _ := strconv.Atoi(slice[i2])
			sum = sum + atoi
		}

		// a以上でb以下の場合はresultに加算
		if a <= sum && sum <= b {
			resultSum = resultSum + i
		}
	}

	fmt.Print(resultSum)
}
