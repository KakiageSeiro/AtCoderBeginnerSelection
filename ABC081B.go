package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ###
// https://atcoder.jp/contests/abs/tasks/abc081_b
// ###

var sc = bufio.NewScanner(os.Stdin)

func main() {
	n, a := nextLine(), nextLine()

	intSlice := createIntSlice(a, n)

	// 割り算を実行した回数
	divisionCount, _ := division(intSlice, 0)

	fmt.Printf("%d", divisionCount)
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func createIntSlice(a string, n string) []int {
	var intSlice []int
	for _, s := range strings.Split(a, " ") {
		atoi, err := strconv.Atoi(s)
		if err != nil {
			panic("Aに整数以外の値が入力されました")
		}

		intSlice = append(intSlice, atoi)
	}

	atoi, _ := strconv.Atoi(n)
	if atoi != len(intSlice) {
		panic("Nの値と、Aの個数が違います")
	}
	return intSlice
}

func division(intSlice []int, divisionCount int) (int, []int) {
	var harfSlice []int //割り算実行後の値を保持するスライス
	for _, i2 := range intSlice {
		if i2%2 != 0 {
			// 奇数が一つでもある場合は処理終了
			return divisionCount, intSlice
		}

		harfSlice = append(harfSlice, i2/2)
	}

	return division(harfSlice, divisionCount+1)
}
