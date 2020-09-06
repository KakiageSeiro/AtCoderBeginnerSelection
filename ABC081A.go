package main

import (
	"fmt"
)

// ###
// https://atcoder.jp/contests/abs/tasks/abc081_a
// ###

func main() {
	var s string
	fmt.Scanf("%s", &s)

	var s1 = s[:1]
	var s2 = s[1:2]
	var s3 = s[2:3]

	const marble = "1"
	var count int
	if s1==marble {
		count++
	}

	if s2==marble {
		count++
	}

	if s3==marble {
		count++
	}

	fmt.Print(count)
}
