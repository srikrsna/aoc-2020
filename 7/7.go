package seven

import (
	"bufio"
	"errors"
	"fmt"
	"strings"

	"github.com/srikrsna/aoc-2020/common"
)

func First() (int, error) {
	bags, err := input()
	if err != nil {
		return 0, err
	}

	mb := bags["shiny gold"]
	holdings := map[string]bool{}
	traverseFirst(bags, mb, holdings)

	return len(holdings), nil
}

func traverseFirst(bags map[string]*bag, cb *bag, holdings map[string]bool) {
	for pn := range cb.Parents {
		if !holdings[pn] {
			holdings[pn] = true
		}

		traverseFirst(bags, bags[pn], holdings)
	}
}

func Second() (int, error) {
	bags, err := input()
	if err != nil {
		return 0, err
	}

	mb := bags["shiny gold"]

	return traverseSecond(bags, mb), nil
}

func traverseSecond(bags map[string]*bag, cb *bag) int {
	sum := 0
	for pn, c := range cb.Contains {
		sum += c + c*traverseSecond(bags, bags[pn])
	}

	return sum
}

type bag struct {
	Name     string
	Parents  map[string]int
	Contains map[string]int
}

func newBag(name string) *bag {
	return &bag{
		Name:     name,
		Parents:  map[string]int{},
		Contains: map[string]int{},
	}
}

func input() (map[string]*bag, error) {
	bags := map[string]*bag{}
	return bags, common.ScanInput(func(scanner *bufio.Scanner) error {
		for scanner.Scan() {
			v := scanner.Text()
			splits := strings.Split(v, " contain ")
			if len(splits) != 2 {
				return errors.New("input error")
			}

			cbn := strings.TrimSuffix(splits[0], " bags")
			cb := bags[cbn]
			if cb == nil {
				cb = newBag(cbn)
				bags[cbn] = cb
			}

			splits = strings.Split(splits[1], ", ")
			for _, s := range splits {
				var (
					c    int
					f, l string
				)
				fmt.Sscanf(strings.TrimRight(s, "bags."), "%d %s %s", &c, &f, &l)
				bn := f + " " + l
				b := bags[bn]
				if b == nil {
					b = newBag(bn)
					bags[bn] = b
				}
				b.Parents[cbn] = c
				cb.Contains[bn] = c
			}
		}

		return nil
	})
}
