package eleven

import (
	"bufio"
	"strings"

	"github.com/srikrsna/aoc-2020/common"
)

func First() (int, error) {
	return solve(4, adjacentCutoff)
}

func Second() (int, error) {
	return solve(5, firstCutoff)
}

func solve(co int, cfn func(r, c, cf int, in []string) bool) (int, error) {
	in, err := input()
	if err != nil {
		return -1, err
	}

	changed := true
	for changed {
		in, changed = run(in, co, cfn)
	}

	ans := 0
	for _, row := range in {
		for _, cell := range row {
			if cell == occupied {
				ans++
			}
		}
	}

	return ans, nil
}

func run(in []string, co int, cfn func(r, c, cf int, in []string) bool) ([]string, bool) {
	out := make([]string, 0, len(in))
	changed := false
	var sb strings.Builder
	for r, row := range in {
		sb.Reset()
		for c, cell := range row {
			switch cell {
			case occupied:
				if cfn(r, c, co, in) {
					sb.WriteRune(free)
					changed = true
				} else {
					sb.WriteRune(occupied)
				}
			case free:
				if cfn(r, c, 1, in) {
					sb.WriteRune(free)
				} else {
					sb.WriteRune(occupied)
					changed = true
				}
			case floor:
				sb.WriteRune(floor)
			}
		}
		out = append(out, sb.String())
	}

	return out, changed
}

func adjacentCutoff(r, c int, cf int, in []string) bool {
	count := 0
	for i := r - 1; i <= r+1; i++ {
		for j := c - 1; j <= c+1; j++ {
			if i == r && j == c {
				continue
			}

			if i < 0 || i >= len(in) || j < 0 || j >= len(in[r]) {
				continue
			}

			if in[i][j] == occupied {
				count++
				if count == cf {
					return true
				}
			}
		}
	}

	return false
}

func firstCutoff(r, c int, cf int, in []string) bool {
	count := 0

	for i := r + 1; i < len(in); i++ {
		if in[i][c] == free {
			break
		} else if in[i][c] == occupied {
			count++
			if count == cf {
				return true
			}
			break
		}
	}

	for i := r - 1; i >= 0; i-- {
		if in[i][c] == free {
			break
		} else if in[i][c] == occupied {
			count++
			if count == cf {
				return true
			}
			break
		}
	}

	for j := c + 1; j < len(in[r]); j++ {
		if in[r][j] == free {
			break
		} else if in[r][j] == occupied {
			count++
			if count == cf {
				return true
			}
			break
		}
	}

	for j := c - 1; j >= 0; j-- {
		if in[r][j] == free {
			break
		} else if in[r][j] == occupied {
			count++
			if count == cf {
				return true
			}
			break
		}
	}

	for i, j := r-1, c-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if in[i][j] == free {
			break
		} else if in[i][j] == occupied {
			count++
			if count == cf {
				return true
			}
			break
		}
	}

	for i, j := r+1, c+1; i < len(in) && j < len(in[i]); i, j = i+1, j+1 {
		if in[i][j] == free {
			break
		} else if in[i][j] == occupied {
			count++
			if count == cf {
				return true
			}
			break
		}
	}

	for i, j := r+1, c-1; i < len(in) && j >= 0; i, j = i+1, j-1 {
		if in[i][j] == free {
			break
		} else if in[i][j] == occupied {
			count++
			if count == cf {
				return true
			}
			break
		}
	}

	for i, j := r-1, c+1; i >= 0 && j < len(in[i]); i, j = i-1, j+1 {
		if in[i][j] == free {
			break
		} else if in[i][j] == occupied {
			count++
			if count == cf {
				return true
			}
			break
		}
	}

	return false
}

const (
	occupied = '#'
	free     = 'L'
	floor    = '.'
)

func input() ([]string, error) {
	var in []string
	return in, common.ScanInput(func(scanner *bufio.Scanner) error {
		for scanner.Scan() {
			in = append(in, scanner.Text())
		}

		return nil
	})
}
