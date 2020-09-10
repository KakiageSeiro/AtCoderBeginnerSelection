package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// ###
// https://atcoder.jp/contests/abs/tasks/abc081_b
// ###

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	n := nextLine2(sc)
	nInt, _ := strconv.Atoi(n)

	aSlice := readA(nInt, sc)
	// 数字が大きい順(降順)に並び替え
	// https://text.baldanders.info/golang/sort/
	sort.Sort(sort.Reverse(sort.IntSlice(aSlice)))

	// アリス、ボブの順で大きい数字を取っていく
	var alice, bob []int
	for i, v := range aSlice {
		if i % 2 == 0 {
			alice = append(alice, v)
		} else {
			bob = append(bob, v)
		}
	}

	// アリスが取った数字の合計
	sumAlice := 0
	for _, v := range alice {
		sumAlice = sumAlice + v
	}

	// ボブが取った数字の合計
	sumBob := 0
	for _, v := range bob {
		sumBob = sumBob + v
	}

	fmt.Print(sumAlice - sumBob)
}

func readA(nInt int, sc *bufio.Scanner) []int{
	var aSlice []int
	a := nextLine2(sc)
	for _, v := range strings.Split(a, " ") {
		vInt, _ := strconv.Atoi(v)
		aSlice = append(aSlice, vInt)
	}

	if len(aSlice) != nInt {
		panic("nの値と、aの数が違います")
	}

	return aSlice
}

func nextLine2(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}
