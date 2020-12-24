package nine

import (
	"bufio"
	"math"
	"strconv"

	"github.com/srikrsna/aoc-2020/common"
)

func First() (int, error) {
	in, err := input()
	if err != nil {
		return -1, err
	}

	return findMismatch(in)
}

func Second() (int, error) {
	in, err := input()
	if err != nil {
		return -1, err
	}

	m, err := findMismatch(in)
	if err != nil {
		return -1, err
	}

	l, r, c := 0, 1, in[0]+in[1]
	for {
		for c < m {
			r++
			c += in[r]
		}

		if c == m {
			var min, max = math.MaxInt64, -1
			for q := l; q <= r; q++ {
				if in[q] < min {
					min = in[q]
				}

				if in[q] > max {
					max = in[q]
				}
			}
			return min + max, nil
		}

		c -= in[r]
		r--

		c -= in[l]
		l++
	}
}

func findMismatch(in []int) (int, error) {
	indexes := map[int][]int{}
	for i, v := range in {
		indexes[v] = append(indexes[v], i)
	}

	const preamble = 25
outer:
	for i, v := range in[preamble:] {
		i := i + preamble
		for j := i - 1; j >= i-preamble; j-- {
			otherHalf := v - in[j]
			for _, ind := range indexes[otherHalf] {
				if ind >= i-preamble && ind <= i-1 {
					continue outer
				}
			}
		}

		return v, nil
	}

	return -1, nil
}

func input() ([]int, error) {
	var ii []int
	return ii, common.ScanInput(func(scanner *bufio.Scanner) error {
		for scanner.Scan() {
			i, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return err
			}
			ii = append(ii, i)
		}

		return nil
	})
}
