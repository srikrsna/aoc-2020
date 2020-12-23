package three

import (
	"bufio"

	"github.com/srikrsna/aoc-2020/common"
)

func First() (int, error) {
	ii, err := input()
	if err != nil {
		return 0, nil
	}

	return solve(ii, 3, 1), nil
}

func Second() (int, error) {
	ii, err := input()
	if err != nil {
		return 0, nil
	}

	return solve(ii, 1, 1) * solve(ii, 3, 1) * solve(ii, 5, 1) * solve(ii, 7, 1) * solve(ii, 1, 2), nil
}

func solve(ii []string, rightStep, downStep int) int {
	var i, j, count int
	for i < len(ii) {
		if ii[i][j] == '#' {
			count++
		}

		j = (j + rightStep) % len(ii[i])
		i += downStep
	}

	return count
}

func input() ([]string, error) {
	var ii []string
	return ii, common.ScanInput(func(scanner *bufio.Scanner) error {
		for scanner.Scan() {
			ii = append(ii, scanner.Text())
		}

		return nil
	})
}
