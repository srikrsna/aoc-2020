package four

import (
	"bufio"
	"strconv"

	"github.com/srikrsna/aoc-2020/common"
)

func First() (int64, error) {
	ii, err := input()
	if err != nil {
		return 0, nil
	}

	count := int64(0)
	for _, p := range ii {
		if p.byr == "" || p.iyr == "" || p.eyr == "" || p.hgt == "" || p.hcl == "" || p.ecl == "" || p.pid == "" {
			continue
		}

		count++
	}

	return count, nil

}

func Second() (int64, error) {
	ii, err := input()
	if err != nil {
		return 0, nil
	}

	count := int64(0)
outer:
	for _, p := range ii {
		if p.byr == "" || p.iyr == "" || p.eyr == "" || p.hgt == "" || p.hcl == "" || p.ecl == "" || p.pid == "" {
			continue
		}

		if len(p.hgt) < 4 || len(p.pid) != 9 || len(p.byr) != 4 || len(p.eyr) != 4 || p.eyr[:2] != "20" || len(p.iyr) != 4 || p.iyr[:2] != "20" || len(p.hcl) != 7 || p.hcl[0] != '#' {
			continue
		}

		if !((p.iyr[2] == '1' && (p.iyr[3] >= '0' && p.iyr[3] <= '9')) || (p.iyr[2] == '2' && p.iyr[3] == '0')) {
			continue
		}

		if !((p.eyr[2] == '2' && (p.eyr[3] >= '0' && p.eyr[3] <= '9')) || (p.eyr[2] == '3' && p.eyr[3] == '0')) {
			continue
		}

		if p.ecl != "amb" && p.ecl != "blu" && p.ecl != "brn" && p.ecl != "gry" && p.ecl != "grn" && p.ecl != "hzl" && p.ecl != "oth" {
			continue
		}

		hu := ""
		hgt := 0
		for i, c := range p.hgt {
			if c < '0' || c > '9' {
				hu = p.hgt[i:]
				break
			}

			hgt *= 10
			hgt += int(c - '0')

			if i > 2 {
				break
			}
		}

		if hu == "cm" {
			if hgt < 150 || hgt > 193 {
				continue
			}
		} else if hu == "in" {
			if hgt < 59 || hgt > 76 {
				continue
			}
		} else {
			continue
		}

		for i := 1; i < len(p.hcl); i++ {
			if !((p.hcl[i] >= '0' && p.hcl[i] <= '9') || (p.hcl[i] >= 'a' && p.hcl[i] <= 'f')) {
				continue outer
			}
		}

		for _, c := range p.pid {
			if c < '0' || c > '9' {
				continue outer
			}
		}

		byr, err := strconv.Atoi(p.byr)
		if err != nil {
			continue
		}
		if byr < 1920 || byr > 2002 {
			continue
		}

		count++
	}

	return count, nil
}

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func input() ([]*passport, error) {
	var ii []*passport
	return ii, common.ScanInput(func(scanner *bufio.Scanner) error {
		i := &passport{}
		for scanner.Scan() {
			v := scanner.Text()
			if v == "" {
				ii = append(ii, i)
				i = &passport{}
			}

			uf := func(key, value string) {
				switch key {
				case "byr":
					i.byr = value
				case "iyr":
					i.iyr = value
				case "eyr":
					i.eyr = value
				case "hgt":
					i.hgt = value
				case "hcl":
					i.hcl = value
				case "ecl":
					i.ecl = value
				case "pid":
					i.pid = value
				case "cid":
					i.cid = value
				}
			}

			colonFound := false
			key, value := "", ""
			for _, c := range v {
				if c == ' ' {
					uf(key, value)
					colonFound = false
					key, value = "", ""
					continue
				}

				if c == ':' {
					colonFound = true
					continue
				}

				if colonFound {
					value += string(c)
				} else {
					key += string(c)
				}
			}
			if colonFound {
				uf(key, value)
			}
		}
		ii = append(ii, i)

		return nil
	})
}
