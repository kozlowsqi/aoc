package main

import (
	"bufio"
	"embed"
	"errors"
	"fmt"
	"io"
	"unicode"
)

//go:embed *.input
var input embed.FS

type item struct {
	value rune
	last  bool
}

type instruction struct {
	from, to, count int
}

func lex(f io.Reader, items chan item, instructions chan instruction) {
	fb := bufio.NewReader(f)

	for {
		crate := make([]byte, 4)

		n, err := fb.Read(crate)
		if err != nil {
			panic(err)
		} else if n != 4 {
			panic("unexpected input length, expected 4 bytes")
		}

		val := rune(crate[1])
		if unicode.IsLetter(val) || val == ' ' {
			items <- item{value: val, last: crate[3] == '\n'}
			continue
		} else if unicode.IsDigit(val) && crate[3] == '\n' {
			break
		}
	}

	close(items)

	for {
		_, err := fb.Discard(1) // discard the newline
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			panic(err)
		}

		var ins instruction
		_, err = fmt.Fscanf(fb, "move %d from %d to %d", &ins.count, &ins.from, &ins.to)
		if err != nil {
			panic(err)
		}

		// correct indexes
		ins.from -= 1
		ins.to -= 1

		instructions <- ins
	}

	close(instructions)
}

func p1(f io.Reader) string {
	items, instructions := make(chan item), make(chan instruction)
	go lex(f, items, instructions)

	crates := make([][]rune, 9)
	var i int
	for it := range items {
		if it.value != ' ' {
			crates[i] = append(crates[i], it.value)
		}

		i++

		if it.last {
			i = 0
		}
	}

	for ins := range instructions {
		for i := 0; i < ins.count; i++ {
			head := crates[ins.from][0]
			crates[ins.to] = append([]rune{head}, crates[ins.to]...)
			crates[ins.from] = crates[ins.from][1:]
		}
	}

	var result []rune
	for _, v := range crates {
		if v != nil {
			result = append(result, v[0])
		}
	}

	return fmt.Sprint(string(result))
}

func p2(f io.Reader) string {
	items, instructions := make(chan item), make(chan instruction)
	go lex(f, items, instructions)

	crates := make([][]rune, 9)
	var i int
	for it := range items {
		if it.value != ' ' {
			crates[i] = append(crates[i], it.value)
		}

		i++

		if it.last {
			i = 0
		}
	}

	for ins := range instructions {
		head := append([]rune{}, crates[ins.from][:ins.count]...)
		crates[ins.to] = append(head, crates[ins.to]...)
		crates[ins.from] = crates[ins.from][ins.count:]
	}

	var result []rune
	for _, v := range crates {
		if v != nil {
			result = append(result, v[0])
		}
	}

	return fmt.Sprint(string(result))
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
