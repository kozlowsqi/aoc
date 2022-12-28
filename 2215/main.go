package main

import (
	"embed"
	"fmt"
	"io"
)

//go:embed *.input
var input embed.FS

func p1(f io.Reader, y int) string {
	rs := rangeset{}
	for {
		var x1, y1, x2, y2 int
		_, err := fmt.Fscanf(f, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &x1, &y1, &x2, &y2)
		if err != nil {
			break
		}

		d := abs(x1-x2) + abs(y1-y2)
		dy := abs(y1 - y)

		if d < dy {
			continue
		}

		rs.add(set{x1 - d + dy, x1 + d - dy})
	}

	return fmt.Sprint(rs.sum())
}

func p2(f io.Reader, m int) string {
	rs := make([]rangeset, m+1)

	for {
		var x1, y1, x2, y2 int
		_, err := fmt.Fscanf(f, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &x1, &y1, &x2, &y2)
		if err != nil {
			break
		}

		d := abs(x1-x2) + abs(y1-y2)
		for y := max(0, y1-d); y <= min(y1+d, m); y++ {
			dy := abs(y1 - y)
			rs[y].add(set{x1 - d + dy, x1 + d - dy})
		}
	}

	for y := range rs {
		rs[y].clip(set{0, m})
		if rs[y].sum() != m && len(rs[y].sets) >= 2 {
			x := (rs[y].sets[0].max + rs[y].sets[1].min) / 2
			return fmt.Sprint(4000000*x + y)
		}
	}

	panic("edge case")
}

func main() {
	f, err := input.Open("test.input")
	die(err)
	fmt.Println(p1(f, 10))

	f, err = input.Open("puzzle.input")
	die(err)
	fmt.Println(p1(f, 2000000))

	f, err = input.Open("test.input")
	die(err)
	fmt.Println(p2(f, 20))

	f, err = input.Open("puzzle.input")
	die(err)
	fmt.Println(p2(f, 4000000))
}

func die(err error) {
	if err != nil {
		panic(err)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
