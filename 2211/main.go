package main

import (
	"bufio"
	"embed"
	"fmt"
	"io"
	"strconv"
	"strings"
	"sort"
	"math/big"
)

//go:embed *.input
var input embed.FS

type monkey struct {
	items   []*big.Int
	op      func(*big.Int)
	test    *big.Int
	onTrue  int64
	onFalse int64
}

func strip(s, prefix string) (int64, bool) {

	if !strings.HasPrefix(s, prefix) {
		return 0, false

	}

	v, err := strconv.Atoi(strings.TrimPrefix(s, prefix))
	die(err)

	return int64(v), true
}

func parse(f io.Reader) []monkey {
	fb := bufio.NewReader(f)

	var m []monkey

	for {
		line, err := fb.ReadString('\n')
		if err == io.EOF {
			break
		}
		die(err)

		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "Monkey") {
			m = append(m, monkey{})

		} else if strings.HasPrefix(line, "Starting items:") {
			line = strings.TrimPrefix(line, "Starting items: ")

			for _, v := range strings.Split(line, ", ") {
				v, err := strconv.Atoi(v)
				die(err)

				m[len(m)-1].items = append(m[len(m)-1].items, big.NewInt(int64(v)))
			}

		} else if strings.HasPrefix(line, "Operation: new = old * old") {
			m[len(m)-1].op = func(a *big.Int) { a.Mul(a, a) }
		} else if n, ok := strip(line, "Operation: new = old * "); ok {
			m[len(m)-1].op = func(a *big.Int) { a.Mul(a, big.NewInt(n)) }
		} else if n, ok := strip(line, "Operation: new = old + "); ok {
			m[len(m)-1].op = func(a *big.Int) { a.Add(a, big.NewInt(n)) }
		} else if n, ok := strip(line, "Test: divisible by "); ok {
			m[len(m)-1].test = big.NewInt(n)
		} else if n, ok := strip(line, "If true: throw to monkey "); ok {
			m[len(m)-1].onTrue = n
		} else if n, ok := strip(line, "If false: throw to monkey "); ok {
			m[len(m)-1].onFalse = n
		}
	}

	return m
}

func p1(f io.Reader) string {
	ms := parse(f)
	count := make([]int, len(ms))

	for k := 0; k < 20; k++ {
		for j := range ms {
			for i := range ms[j].items {
				ms[j].op(ms[j].items[i])
				ms[j].items[i].Quo(ms[j].items[i], big.NewInt(3))
				count[j]++

				t := big.NewInt(0).Set(ms[j].items[i])
				if t.Mod(ms[j].items[i], ms[j].test).Cmp(big.NewInt(0)) == 0 {
					ms[ms[j].onTrue].items = append(ms[ms[j].onTrue].items, ms[j].items[i])
				} else {
					ms[ms[j].onFalse].items = append(ms[ms[j].onFalse].items, ms[j].items[i])
				}
			}

			ms[j].items = nil
		}
	}
	sort.Ints(count)
	return fmt.Sprint(count[len(count)-1]*count[len(count)-2])
}

func p2(f io.Reader) string {
	ms := parse(f)
	count := make([]int, len(ms))

	lcm := big.NewInt(1)
	for _, m := range ms {
		gcm := big.NewInt(0).GCD(nil, nil, lcm, m.test)
		lcm.Set(gcm.Div(lcm, gcm).Mul(gcm, m.test))
	}

	for k := 0; k < 10000; k++ {
		for j := range ms {
			for i := range ms[j].items {
				ms[j].op(ms[j].items[i])
				ms[j].items[i].Mod(ms[j].items[i], lcm)
				count[j]++

				t := big.NewInt(0).Set(ms[j].items[i])
				if t.Mod(ms[j].items[i], ms[j].test).Cmp(big.NewInt(0)) == 0 {
					ms[ms[j].onTrue].items = append(ms[ms[j].onTrue].items, ms[j].items[i])
				} else {
					ms[ms[j].onFalse].items = append(ms[ms[j].onFalse].items, ms[j].items[i])
				}
			}

			ms[j].items = nil
		}
	}
	sort.Ints(count)
	return fmt.Sprint(count[len(count)-1]*count[len(count)-2])
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
