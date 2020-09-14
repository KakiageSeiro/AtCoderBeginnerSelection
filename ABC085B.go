package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// ###
// https://atcoder.jp/contests/abs/tasks/abc085_b
// ###

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	n := nextLine3(sc)
	nInt, _ := strconv.Atoi(n)

	mochiDiameterSlice := readMochiDiameter(nInt, sc)

	// 数字(直径)が大きい順(降順)に並び替え
	// https://text.baldanders.info/golang/sort/
	sort.Sort(sort.Reverse(sort.IntSlice(mochiDiameterSlice)))

	dan := 0 //鏡餅の段数
	beforeMochiDiameter := 101 //前回ループ時の直径。入力値の制約で100より大きい値が入力されることはない
	for _, v := range mochiDiameterSlice {
		if beforeMochiDiameter <= v {
			// 下の段(前回のループ時の直径)以上の餅は乗せられない
			continue
		}

		beforeMochiDiameter = v
		dan++
	}

	fmt.Print(dan)
}

func readMochiDiameter(n int, sc *bufio.Scanner) []int{
	var mochiDiameterSlice []int

	for i := 0; i < n; i++ {
		d := nextLine3(sc)
		dInt, _ := strconv.Atoi(d)
		mochiDiameterSlice = append(mochiDiameterSlice, dInt)
	}

	if len(mochiDiameterSlice) != n {
		panic("nの値と、dの数が違います")
	}

	return mochiDiameterSlice
}

func nextLine3(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}
