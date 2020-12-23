package six

import (
	"bufio"

	"github.com/srikrsna/aoc-2020/common"
)

func First() (int, error) {
	return solve(func(c int, _ []string) bool {
		return c > 0
	})
}

func Second() (int, error) {
	return solve(func(c int, grp []string) bool {
		return c == len(grp)
	})
}

func solve(cf func(int, []string) bool) (int, error) {
	ii, err := input()
	if err != nil {
		return 0, nil
	}

	sum := 0
	for _, i := range ii {
		sum += yesCount(i, cf)
	}

	return sum, nil
}

func yesCount(grp []string, cf func(int, []string) bool) int {
	alps := [26]int{}
	for _, s := range grp {
		for _, c := range s {
			alps[c-'a']++
		}
	}

	count := 0
	for _, c := range alps {
		if cf(c, grp) {
			count++
		}
	}

	return count
}

func input() ([][]string, error) {
	var out [][]string
	return out, common.ScanInput(func(scanner *bufio.Scanner) error {
		var ii []string
		for scanner.Scan() {
			v := scanner.Text()
			if v == "" {
				out = append(out, ii)
				ii = nil
				continue
			}
			ii = append(ii, v)
		}
		out = append(out, ii)

		return nil
	})
}
