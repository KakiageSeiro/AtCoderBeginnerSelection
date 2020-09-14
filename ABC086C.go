package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// ###
// https://atcoder.jp/contests/abs/tasks/arc089_a
// ###

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	n := nextLineABC086C(sc)
	nInt, _ := strconv.Atoi(n)

	planSlice := readABC086C(nInt, sc)

	t := 0
	x := 0
	y := 0
	for _, p := range planSlice {
		diffT := p.t - t
		diffX := float64(p.x - x)
		diffY := float64(p.y - y)

		f := int(math.Abs(diffX) + math.Abs(diffY))

		if diffT < f || diffT%2 != f%2 {
			fmt.Print("No")
			return
		}

		t = p.t
		x = p.x
		y = p.y
	}

	fmt.Print("Yes")
}

func readABC086C(nInt int, sc *bufio.Scanner) []plan {
	var planSlice []plan

	for i := 0; i < nInt; i++ {

		var t, x, y int
		split := strings.Split(nextLineABC086C(sc), " ")

		t, _ = strconv.Atoi(split[0])
		x, _ = strconv.Atoi(split[1])
		y, _ = strconv.Atoi(split[2])

		p := plan{
			t: t,
			x: x,
			y: y,
		}

		planSlice = append(planSlice, p)
	}

	if len(planSlice) != nInt {
		panic("nの値と、aの数が違います")
	}

	return planSlice
}

type plan struct {
	t int
	x int
	y int
}

func nextLineABC086C(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}
