package eight

import (
	"bufio"
	"fmt"

	"github.com/srikrsna/aoc-2020/common"
)

func First() (int, error) {
	ins, err := input()
	if err != nil {
		return -1, nil
	}

	visited := make([]bool, len(ins))
	res := 0
	for i := 0; i < len(ins); i++ {
		if visited[i] {
			return res, nil
		}

		visited[i] = true

		switch ins[i].Cmd {
		case acc:
			res += ins[i].Op
		case jmp:
			i += ins[i].Op
			i--
		case nop:
		}
	}

	return -1, nil
}

func Second() (int, error) {
	ins, err := input()
	if err != nil {
		return -1, nil
	}

	visited := make([]bool, len(ins))
	res := 0
	executions := make([]*step, 0, len(ins))
	var changeStep *step
	for i := 0; i < len(ins); i++ {
		if visited[i] {
			// Unwind
			j := len(executions) - 1
			if changeStep != nil {
				j = changeStep.es - 1
			}
			for ; j >= 0; j-- {
				ind := executions[j].i
				visited[ind] = false

				cmd := ins[ind].Cmd
				if cmd == nop || cmd == jmp {
					changeStep = executions[j]
					break
				}
			}
			executions = executions[:j]
			res = changeStep.acc
			i = changeStep.i - 1

			continue
		}

		visited[i] = true

		executions = append(executions, &step{
			i, len(executions), res,
		})
		switch ins[i].Cmd {
		case acc:
			res += ins[i].Op
		case jmp:
			if changeStep == nil || changeStep.i != i {
				i += ins[i].Op
				i--
			}
		case nop:
			if changeStep != nil && changeStep.i == i {
				i += ins[i].Op
				i--
			}
		}

	}

	return res, nil
}

type step struct {
	i   int
	es  int
	acc int
}

const (
	acc = "acc"
	jmp = "jmp"
	nop = "nop"
)

type ins struct {
	Cmd string
	Op  int
}

func input() ([]*ins, error) {
	var ii []*ins
	return ii, common.ScanInput(func(scanner *bufio.Scanner) error {
		for scanner.Scan() {
			var i ins
			if _, err := fmt.Sscanf(scanner.Text(), "%s %d", &i.Cmd, &i.Op); err != nil {
				return err
			}
			ii = append(ii, &i)
		}

		return nil
	})
}
