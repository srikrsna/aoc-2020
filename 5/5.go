package five

import (
	"bufio"

	"github.com/srikrsna/aoc-2020/common"
)

func First() (int, error) {
	ii, err := input()
	if err != nil {
		return 0, nil
	}

	maxId := -1
	for _, i := range ii {
		row := solve(i[:7], 127, 2, 'F', 'B')
		col := solve(i[7:], 7, 2, 'L', 'R')
		id := row*8 + col
		if id > maxId {
			maxId = id
		}
	}

	return maxId, nil
}

func Second() (int, error) {
	ii, err := input()
	if err != nil {
		return 0, nil
	}

	ids := [8][128]int{}

	for _, i := range ii {
		row := solve(i[:7], 127, 2, 'F', 'B')
		col := solve(i[7:], 7, 2, 'L', 'R')
		ids[col][row] = row*8 + col
	}

	for c, col := range ids {
		for r, cid := range col {
			if r == 0 || r == 127 {
				continue
			}

			// Present so continue
			if cid > 0 {
				continue
			}

			id := r*8 + c

			if checkSeat(ids, id-1) && checkSeat(ids, id+1) {
				return id, nil
			}
		}
	}

	return -1, nil
}

func checkSeat(grid [8][128]int, id int) bool {
	for i := 0; i < 8; i++ {
		if (id-i)%8 != 0 {
			continue
		}

		r := (id - i) / 8
		if r > 127 || r < 0 {
			continue
		}

		if grid[i][r] > 0 {
			return true
		}
	}

	return false
}

func solve(code string, max, p int, f, s rune) int {
	low, high := 0, max
	for _, c := range code {
		mid := (high - low) / p
		if c == s {
			low += mid + 1
		} else {
			high = low + mid
		}
	}

	return low
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
