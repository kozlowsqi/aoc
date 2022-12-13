package main

import (
	"embed"
	"errors"
	"fmt"
	"io"
	"sort"
)

//go:embed *.input
var input embed.FS

func p1(f io.Reader) string {
	var sum, max int

	for {
		var x int
		n, err := fmt.Fscanln(f, &x)

		if n == 1 {
			sum += x
			continue
		}

		if sum > max {
			max = sum
		}

		if n == 0 {
			sum = 0
		}

		if errors.Is(err, io.EOF) {
			return fmt.Sprint(max)
		}
	}
}

func p2(f io.Reader) string {
	var (
		sum int
		sums []int
	)

	for {
		var x int
		n, err := fmt.Fscanln(f, &x)

		if n == 1 {
			sum += x
			continue
		}

		if n == 0 {
			sums = append(sums, sum)
			sum = 0
		}

		if errors.Is(err, io.EOF) {
			sort.Ints(sums)
			l := len(sums) - 1
			return fmt.Sprint(sums[l] + sums[l-1] + sums[l-2])
		}
	}
}

func main() {
	f, err := input.Open("test.input")
	if err != nil {
		panic(err)
	}

	fmt.Println(p1(f))

	f, err = input.Open("puzzle.input")
	if err != nil {
		panic(err)
	}

	fmt.Println(p1(f))

	f, err = input.Open("test.input")
	if err != nil {
		panic(err)
	}

	fmt.Println(p2(f))

	f, err = input.Open("puzzle.input")
	if err != nil {
		panic(err)
	}

	fmt.Println(p2(f))
}
