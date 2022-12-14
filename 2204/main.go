package main

import (
	"embed"
	"fmt"
	"io"
)

//go:embed *.input
var input embed.FS


func p1(f io.Reader) string {
	var total int
	for {
		var x1, x2, y1, y2 int
		_, err := fmt.Fscanf(f, "%d-%d,%d-%d", &x1, &x2, &y1, &y2)
		if err == io.EOF {
			return fmt.Sprint(total)
		}
		die(err)

		if x1 >= y1 && x2 <= y2 || x1 <= y1 && x2 >= y2 {
			total++
		}
	}
}

func p2(f io.Reader) string {
	var total int
	for {
		var x1, x2, y1, y2 int
		_, err := fmt.Fscanf(f, "%d-%d,%d-%d", &x1, &x2, &y1, &y2)
		if err == io.EOF {
			return fmt.Sprint(total)
		}
		die(err)

		if x1 <= y2 && y1 <= x2 {
			total++
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

