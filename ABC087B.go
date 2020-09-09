package main

import (
	"fmt"
)

// ###
// https://atcoder.jp/contests/abs/tasks/abc087_b
// ###

func main() {

	var a, b, c, x int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	fmt.Scanf("%d", &x)

	jastCount := createChooseSlice(a, b, c, x)

	fmt.Print(jastCount)
}

// コイン選択方法の組み合わせ
func createChooseSlice(a, b, c, xYen int) int {
	jastCount := 0
	for i := 0; i < a+1; i++ { // 500のコインが0枚の場合でも、0枚を選択する組み合わせを実行したいため、条件式に+1を指定(iが0のときはループを実行される)
		for i2 := 0; i2 < b+1; i2++ {
			for i3 := 0; i3 < c+1; i3++ {
				sum := (i * 500) + (i2 * 100) + (i3 * 50)

				if sum == xYen {
					jastCount++
				}
			}
		}
	}

	return jastCount
}