package main

import (
	"embed"
	"fmt"
	"io"
)

//go:embed *.input
var input embed.FS

func set(s string) []rune {
	set := make([]rune, 53)

	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			set[int(r) - int('A') + 27]++
		} else if r <= 'z' && r >= 'a' {
			set[int(r) - int('a') + 1]++
		} else {
			panic("out of range")
		}
	}

	return set
}

func intersect(a, b []rune) []rune {
	set := make([]rune, 53)

	for i := range set {
		if a[i] != 0 && b[i] != 0 {
			set[i]++
		}
	}

	return set
}

func p1(f io.Reader) string {
	var total int
	for {
		var x string
		_, err := fmt.Fscanln(f, &x)
		if err == io.EOF {
			return fmt.Sprint(total)
		} else {
			die(err)
		}

		x1, x2 := set(x[:len(x)/2]), set(x[len(x)/2:])

		for i, v := range intersect(x1, x2) {
			if v != 0 { total += i }
		}
	}
}

func p2(f io.Reader) string {
	var total int
	for {
		var x string
		_, err := fmt.Fscanln(f, &x)
		if err == io.EOF {
			return fmt.Sprint(total)
		} else {
			die(err)
		}

		s := set(x)

		_, err = fmt.Fscanln(f, &x)
		if err == io.EOF {
			return fmt.Sprint(total)
		} else {
			die(err)
		}

		s = intersect(s, set(x))

		_, err = fmt.Fscanln(f, &x)
		if err == io.EOF {
			return fmt.Sprint(total)
		} else {
			die(err)
		}

		for i, v := range intersect(s, set(x)) {
			if v != 0 { total += i }
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
	if err != nil { panic(err) }
}

