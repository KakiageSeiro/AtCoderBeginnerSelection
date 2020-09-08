package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ###
// https://atcoder.jp/contests/abs/tasks/abc087_b
// ###

// 標準入力から一行読み取る
func nextLine_abc087_b() string {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()
}

func main() {
	a, b, c, x := nextLine_abc087_b(), nextLine_abc087_b(), nextLine_abc087_b(), nextLine_abc087_b()
	coinCount := cast(a, b, c)
	xYen, _ := strconv.Atoi(x)

	chooses := createChooseSlice(coinCount)

	// 指定された合計金額と一致するコインの組み合わせ数
	jastCount := 0
	for _, ch := range chooses {
		sum := (ch.count500 * 500) + (ch.count100 * 100) + (ch.count50 * 50)

		if sum == xYen {
			jastCount++
		}
	}

	fmt.Print(jastCount)
}

// コイン選択方法の組み合わせ
func createChooseSlice(coinCount coinCount) []choose {
	chooses := []choose{}
	for i := 0; i < coinCount.a500+1; i++ { // 500のコインが0枚の場合でも、0枚を選択する組み合わせを実行したいため、条件式に+1を指定(iが0のときはループを実行される)
		currentChoose := choose{}
		currentChoose.count500 = i

		for i := 0; i < coinCount.b100+1; i++ {
			currentChoose.count100 = i

			for i := 0; i < coinCount.c50+1; i++ {
				currentChoose.count50 = i

				chooses = append(chooses, currentChoose)
			}
		}
	}

	return chooses
}

func cast(a, b, c string) coinCount {
	a500, _ := strconv.Atoi(a)
	b100, _ := strconv.Atoi(b)
	c50, _ := strconv.Atoi(c)

	return coinCount{
		a500: a500,
		b100: b100,
		c50:  c50,
	}
}

// 入力されたコインの枚数
type coinCount struct {
	a500 int
	b100 int
	c50  int
}

// 各コインの選び方の組み合わせ
type choose struct {
	count500 int
	count100 int
	count50  int
}
