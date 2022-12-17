package main

import (
	"bufio"
	"embed"
	"fmt"
	"io"
)

//go:embed *.input
var input embed.FS

func p1(f io.Reader) string {
	fb := bufio.NewReader(f)

	var c1, c2, c3 rune
	for i := 1; ; i++ {
		r, _, err := fb.ReadRune()
		die(err)

		if c1 > 0 && c1 != c2 && c1 != c3 && c1 != r && c2 != c3 && c2 != r && c3 != r {
			return fmt.Sprint(i)
		}

		c1, c2, c3 = c2, c3, r
	}
}

func p2(f io.Reader) string {
	fb := bufio.NewReader(f)

	cMap := map[rune]int{}
	cQueue := []rune{}
	for i := 1; ; i++ {
		r, _, err := fb.ReadRune()
		die(err)
		cMap[r]++

		cQueue = append(cQueue, r)
		if len(cQueue) > 14 {
			cMap[cQueue[0]]--
			cQueue = cQueue[1:]
		}

		u := true
		for _, v := range cMap {
			if v > 1 {
				u = false
			}
		}

		if u && len(cQueue) == 14 {
			return fmt.Sprint(i)
		}
	}
}

func main() {
	f, err := input.Open("test.input")
	die(err)
	fmt.Println(p1(f))

	f, err = input.Open("puzzle.input")
	die(err)
	fmt.Println(p1(f))

	f, err = input.Open("test.input")
	die(err)
	fmt.Println(p2(f))

	f, err = input.Open("puzzle.input")
	die(err)
	fmt.Println(p2(f))
}

func die(err error) {
	if err != nil {
		panic(err)
	}
}
