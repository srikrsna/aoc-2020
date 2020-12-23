package two

import (
	"bufio"
	"fmt"

	"github.com/srikrsna/aoc-2020/common"
)

func First() (int64, error) {
	ii, err := input()
	if err != nil {
		return 0, nil
	}

	count := int64(0)
	for _, i := range ii {
		ccount := int64(0)
		for _, c := range i.Password {
			if c == i.Character {
				ccount++
			}
		}

		if ccount >= i.Min && ccount <= i.Max {
			count++
		}
	}

	return count, nil
}

func Second() (int64, error) {
	ii, err := input()
	if err != nil {
		return 0, nil
	}

	count := int64(0)
	for _, i := range ii {
		ccount := int64(0)
		if i.Password[i.Min-1] == byte(i.Character) {
			ccount++
		}

		if i.Password[i.Max-1] == byte(i.Character) {
			ccount++
		}

		if ccount == 1 {
			count++
		}
	}

	return count, nil
}

type inputLine struct {
	Min, Max  int64
	Character rune
	Password  string
}

func input() ([]*inputLine, error) {
	var ii []*inputLine
	return ii, common.ScanInput(func(scanner *bufio.Scanner) error {
		for scanner.Scan() {
			var i inputLine
			if _, err := fmt.Sscanf(scanner.Text(), "%d-%d %c: %s", &i.Min, &i.Max, &i.Character, &i.Password); err != nil {
				return err
			}

			ii = append(ii, &i)
		}

		return nil
	})
}
