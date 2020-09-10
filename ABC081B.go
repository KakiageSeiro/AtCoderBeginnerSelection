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


func main() {
	n, a := nextLine(), nextLine()

	intslice := createIntSlice(a, n)

	// 割り算を実行した回数
	divisionCount, _ := division(intslice, 0)

	fmt.Printf("%d", divisionCount)
}

func nextLine() string {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()
}

func createIntSlice(a string, n string) []int {
	var intslice []int
	for _, s := range strings.Split(a, " ") {
		atoi, err := strconv.Atoi(s)
		if err != nil {
			panic("Aに整数以外の値が入力されました")
		}

		intslice = append(intslice, atoi)
	}

	atoi, _ := strconv.Atoi(n)
	if atoi != len(intslice) {
		panic("Nの値と、Aの個数が違います")
	}
	return intslice
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