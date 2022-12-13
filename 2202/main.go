package main

import (
	"embed"
	"fmt"
	"io"
)

//go:embed *.input
var input embed.FS

// A, X - rock
// B, Y - paper
// C, Z - scissors

func p1(f io.Reader) string {
	var total int

	for {
		var score int
		var v1, v2 string
		_, err := fmt.Fscanln(f, &v1, &v2)
		if err == io.EOF {
			return fmt.Sprint(total)
		} else {
			die(err)
		}

		switch {
		case v1 == "A" && v2 == "Y", v1 == "B" && v2 == "Z",
			v1 == "C" && v2 == "X":
			score += 6
		case v1 == "A" && v2 == "X", v1 == "B" && v2 == "Y",
			v1 == "C" && v2 == "Z":
			score += 3
		}

		switch v2 {
		case "X":
			score += 1
		case "Y":
			score += 2
		case "Z":
			score += 3
		}

		total += score
	}
}

// A - rock
// B - paper
// C - scissors
//
func p2(f io.Reader) string {
	var total int

	for {
		var score int
		var v1, v2 string
		_, err := fmt.Fscanln(f, &v1, &v2)
		if err == io.EOF {
			return fmt.Sprint(total)
		} else {
			die(err)
		}

		switch {
			// rock 
			case v2 == "X" && v1 == "B", v2 == "Y" && v1 == "A", v2 == "Z" && v1 == "C":
			score += 1
			// paper 
			case v2 == "X" && v1 == "C", v2 == "Y" && v1 == "B", v2 == "Z" && v1 == "A":
			score += 2
			// scissors 
			case v2 == "X" && v1 == "A", v2 == "Y" && v1 == "C", v2 == "Z" && v1 == "B":
			score += 3
		}

		switch v2 {
		case "Y":
			score += 3
		case "Z":
			score += 6
		}

		total += score
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

