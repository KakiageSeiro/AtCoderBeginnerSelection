package main

import (
	"fmt"
	"strings"
)

// ###
// https://atcoder.jp/contests/abs/tasks/arc065_a
// ###

func main() {

	var s string
	fmt.Scanf("%s", &s)

	for {
		if strings.HasSuffix(s, "dreamer") {
			s = s[:len(s) - 7]
		} else if strings.HasSuffix(s, "dream") {
			s = s[:len(s) - 5]
		} else if strings.HasSuffix(s, "eraser") {
			s = s[:len(s) - 6]
		} else if strings.HasSuffix(s, "erase") {
			s = s[:len(s) - 5]
		} else {
			if len(s) == 0 {
				fmt.Print("YES")
			} else {
				fmt.Print("NO")
			}

			return
		}
	}
}
