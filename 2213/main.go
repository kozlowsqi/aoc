package main

import (
	"bufio"
	"embed"
	"fmt"
	"io"
	"strconv"
)

//go:embed *.input
var input embed.FS

func scanList(a []byte) ([]byte, []byte) {
	i := 1
	for n := 1; n != 0 && i < len(a); i++ {
		switch {
		case a[i] == '[':
			n++
		case a[i] == ']':
			n--
		}
	}

	return a[:i], a[i:]
}

func scanNumber(a []byte) ([]byte, []byte) {
	i := 0
	for ; i < len(a) && a[i] != ','; i++ {
	}
	return a[:i], a[i:]
}

func scanValue(a []byte) (al, at []byte) {
	if a[0] == '[' {
		al, at = scanList(a)
	} else {
		al, at = scanNumber(a)
	}

	if len(at) > 0 && at[0] == ',' {
		at = at[1:]
	}

	return
}

func makeList(a []byte) []byte {
	b := make([]byte, len(a)+2)
	b[0], b[len(b)-1] = '[', ']'
	copy(b[1:], a)
	return b
}

type result int

const (
	ro result = iota // right order
	wo               // wrong order
	co               // continue
)

func compareValues(a, b []byte) result {

	if a[0] != '[' && b[0] != '[' {

		n1, err := strconv.Atoi(string(a))
		die(err)

		n2, err := strconv.Atoi(string(b))
		die(err)

		if n1 < n2 {
			return ro
		}

		if n1 > n2 {
			return wo
		}

		return co
	}

	if a[0] != '[' {
		a = makeList(a)
	}

	if b[0] != '[' {
		b = makeList(b)
	}

	return compareLists(a, b)
}

func compareLists(a, b []byte) result {
	a, b = a[1:len(a)-1], b[1:len(b)-1]

	for {
		if len(a) == 0 && len(b) == 0 {
			return co
		}

		if len(a) == 0 && len(b) > 0 {
			return ro
		}

		if len(b) == 0 && len(a) > 0 {
			return wo
		}

		al, at := scanValue(a)
		bl, bt := scanValue(b)

		r := compareValues(al, bl)
		if r != co {
			return r
		}

		a = at
		b = bt
	}
}

func p1(f io.Reader) string {
	fb := bufio.NewReader(f)

	var i, total = 1, 0
	for {
		a, err := fb.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		die(err)

		b, err := fb.ReadBytes('\n')
		die(err)

		if compareLists(a[:len(a)-1], b[:len(b)-1]) == ro {
			total += i
		}

		fb.Discard(1)
		i++
	}

	return fmt.Sprint(total)
}

func mergesort(p [][]byte) [][]byte {
	if len(p) == 1 {
		return p
	}

	h := mergesort(p[:len(p)/2])
	t := mergesort(p[len(p)/2:])

	var s [][]byte
	for len(h) > 0 && len(t) > 0 {
		switch compareLists(h[0], t[0]) {
		case ro:
			s = append(s, h[0])
			h = h[1:]
		case wo:
			s = append(s, t[0])
			t = t[1:]
		}

	}

	if len(h) > 0 {
		s = append(s, h...)
	}

	if len(t) > 0 {
		s = append(s, t...)
	}

	return s
}

func p2(f io.Reader) string {
	fb := bufio.NewReader(f)

	var p [][]byte
	for {
		a, err := fb.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		die(err)

		if len(a) == 1 {
			continue
		}

		a = a[:len(a)-1]
		p = append(p, a)
	}

	p = append(p, []byte("[[2]]"), []byte("[[6]]"))
	sorted := mergesort(p)

	var x, y int
	for i, a := range sorted {
		if string(a) == "[[2]]" {
			x = i + 1
		}

		if string(a) == "[[6]]" {
			y = i + 1
		}
	}

	return fmt.Sprint(x * y)
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
