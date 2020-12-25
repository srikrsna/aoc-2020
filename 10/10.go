package ten

import (
	"bufio"
	"sort"
	"strconv"

	"github.com/srikrsna/aoc-2020/common"
)

func First() (int, error) {
	in, err := input()
	if err != nil {
		return -1, err
	}

	var p, ones, twos, threes int
	for _, v := range in[1:] {
		switch v - p {
		case 1:
			ones++
		case 3:
			threes++
		case 2:
			twos++
		}
		p = v
	}

	return ones * threes, nil
}

func Second() (int, error) {
	in, err := input()
	if err != nil {
		return -1, err
	}

	dp := make([]int, 0, len(in))
	dp = append(dp, 1)
	for i := 1; i < len(in); i++ {
		dp = append(dp, 0)
		for j := i - 1; j >= 0; j-- {
			if in[i]-in[j] > 3 {
				break
			}

			dp[i] += dp[j]
		}
	}

	return dp[len(dp)-1], nil
}

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func input() (h intHeap, err error) {
	h = append(h, 0)
	defer func() {
		sort.Sort(h)
		h = append(h, h[len(h)-1]+3)
	}()
	return h, common.ScanInput(func(scanner *bufio.Scanner) error {
		for scanner.Scan() {
			i, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return err
			}
			h = append(h, i)
		}

		return nil
	})
}
