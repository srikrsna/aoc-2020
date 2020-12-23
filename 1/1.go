package one

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func First() (int64, error) {
	ii, err := input()
	if err != nil {
		return 0, err
	}

	return solve(ii, 2020), nil
}

func Second() (int64, error) {
	ii, err := input()
	if err != nil {
		return 0, err
	}

	for _, i := range ii {
		marker := 2020 - i
		p := solve(ii, marker)
		if p > 0 {
			return p * i, nil
		}
	}

	return -1, nil
}

func solve(ii []int64, marker int64) int64 {
	lookup := map[int64]bool{}
	for _, j := range ii {
		otherSide := marker - j
		lookup[j] = true
		if lookup[otherSide] {
			return j * otherSide
		}
	}

	return -1
}

func input() ([]int64, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("unable to load input: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	input := []int64{}
	for scanner.Scan() {
		curr, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("malformed input: %w", err)
		}

		input = append(input, curr)
	}

	return input, nil
}
