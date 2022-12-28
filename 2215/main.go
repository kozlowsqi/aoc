package main

import (
	"embed"
	"fmt"
	"io"
)

//go:embed *.input
var input embed.FS

type sensor struct {
	x1, y1 int
	x2, y2 int
}

func (s *sensor) d() int {
	return abs(s.x1-s.x2) + abs(s.y1-s.y2)
}

func (s *sensor) dy(y int) int {
	return abs(s.y1 - y)
}

func (s *sensor) dydx(y int) int {
	return s.d() - s.dy(y)
}

func (s *sensor) in(y int) bool {
	return s.dy(y) <= s.d()
}

func p1(f io.Reader, y int) string {
	rs := rangeset{}
	for {
		var s sensor
		_, err := fmt.Fscanf(f, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.x1, &s.y1, &s.x2, &s.y2)
		if err != nil {
			break
		}

		if !s.in(y) {
			continue
		}

		rs.add(set{s.x1 - s.dydx(y), s.x1 + s.dydx(y)})
	}

	return fmt.Sprint(rs.sum())
}

func p2(f io.Reader, m int) string {
	rs := make([]rangeset, m+1)

	for {
		var s sensor
		_, err := fmt.Fscanf(f, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.x1, &s.y1, &s.x2, &s.y2)
		if err != nil {
			break
		}

		for y := max(0, s.y1-s.d()); y <= min(s.y1+s.d(), m); y++ {
			rs[y].add(set{s.x1 - s.dydx(y), s.x1 + s.dydx(y)})
		}
	}

	for y := range rs {
		rs[y].clip(set{0, m})
		if rs[y].sum() != m && len(rs[y].sets) >= 2 {
			x := (rs[y].sets[0].max + rs[y].sets[1].min) / 2
			return fmt.Sprint(4000000*x + y)
		}
	}

	panic("yeah, there are potential edge cases this doesn't cover...")
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
