package main

import (
	"embed"
	"fmt"
	"io"
)

//go:embed *.input
var input embed.FS

func score(y, x int, atrees [][]rune) int {

	var a int
	for i := x - 1; i >= 0; i-- {
		a++

		if atrees[y][i] >= atrees[y][x] {
			break
		}
	}

	var b int
	for i := x + 1; i < len(atrees[y]); i++ {
		b++

		if atrees[y][i] >= atrees[y][x] {
			break
		}
	}

	var c int
	for i := y - 1; i >= 0; i-- {
		c++

		if atrees[i][x] >= atrees[y][x] {
			break
		}
	}

	var d int
	for i := y + 1; i < len(atrees); i++ {
		d++

		if atrees[i][x] >= atrees[y][x] {
			break
		}
	}

	return a * b * c * d
}

func visible(y, x int, atrees [][]rune) bool {
	o := false
	for i := 0; i < x; i++ {
		o = o || atrees[y][i] >= atrees[y][x]
	}
	if !o {
		return true
	}

	o = false
	for i := x + 1; i < len(atrees[y]); i++ {
		o = o || atrees[y][i] >= atrees[y][x]
	}
	if !o {
		return true
	}

	o = false
	for i := 0; i < y; i++ {
		o = o || atrees[i][x] >= atrees[y][x]
	}
	if !o {
		return true
	}

	o = false
	for i := y + 1; i < len(atrees); i++ {
		o = o || atrees[i][x] >= atrees[y][x]
	}

	return !o
}

func p1(f io.Reader) string {
	var atrees [][]rune

	for {
		var s string
		_, err := fmt.Fscanln(f, &s)
		if err == io.EOF {
			break
		}

		die(err)
		atrees = append(atrees, []rune(s))
	}

	var total int
	for i := 0; i < len(atrees); i++ {
		for j := 0; j < len(atrees[i]); j++ {
			if visible(i, j, atrees) {
				total++
			}
		}
	}

	return fmt.Sprint(total)
}

func p2(f io.Reader) string {
	var atrees [][]rune

	for {
		var s string
		_, err := fmt.Fscanln(f, &s)
		if err == io.EOF {
			break
		}

		die(err)
		atrees = append(atrees, []rune(s))
	}

	var total int
	for i := 0; i < len(atrees); i++ {
		for j := 0; j < len(atrees[i]); j++ {
			total = max(score(i, j, atrees), total)
		}
	}

	return fmt.Sprint(total)
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
