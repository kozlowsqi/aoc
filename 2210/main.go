package main

import (
	"embed"
	"fmt"
	"io"
)

//go:embed *.input
var input embed.FS

func p1(f io.Reader) string {
	var clk, x, total int = 0, 1, 0
	for {

		var (
			val int
			cmd string
		)
		_, err := fmt.Fscanln(f, &cmd, &val)
		if err == io.EOF {
			break
		} else if cmd != "noop" {
			die(err)
		}

		switch cmd {
		case "noop":
			clk += 1
			if clk >= 20 && (clk-20)%40 == 0 {
				total += x * clk
			}

		case "addx":
			for i := 0; i < 2; i++ {
				clk += 1

				if clk >= 20 && (clk-20)%40 == 0 {
					total += x * clk
				}
			}

			x += val
		}

	}

	return fmt.Sprint(total)
}

func p2(f io.Reader) string {
	var clk, x = 0, 1
	var result []byte

	for {

		var (
			val int
			cmd string
		)
		_, err := fmt.Fscanln(f, &cmd, &val)
		if err == io.EOF {
			break
		} else if cmd != "noop" {
			die(err)
		}

		switch cmd {
		case "noop":
			clk += 1

			if clk%40 >= x && clk%40 < x+3 {
				result = fmt.Append(result, "█")
			} else {
				result = fmt.Append(result, " ")
			}

			if clk%40 == 0 {
				result = fmt.Appendln(result)
			}

		case "addx":
			for i := 0; i < 2; i++ {
				clk += 1

				if clk%40 >= x && clk%40 < x+3 {
					result = fmt.Append(result, "█")
				} else {
					result = fmt.Append(result, " ")
				}

				if clk%40 == 0 {
					result = fmt.Appendln(result)
				}
			}

			x += val
		}

	}

	return string(result)
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
