package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func fileScanner() *bufio.Scanner {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewScanner(file)
}

type report struct {
	levels []int32
	safe   bool
}

func remove(slice []int32, s int) []int32 {
	return append(slice[:s], slice[s+1:]...)
}

func (r *report) eval() (int, bool) {

	dir := ""

	r.safe = true
	for i, lvl := range r.levels {
		if i == 0 {
			continue
		}

		// If current bigger than last
		if lvl > r.levels[i-1] {
			if dir == "decreasing" {
				r.safe = false
				return i, false
			} else {
				dir = "increasing"
				if lvl-r.levels[i-1] < 1 || lvl-r.levels[i-1] > 3 {
					r.safe = false
					return i, false
				}
			}
		}

		if lvl < r.levels[i-1] {
			if dir == "increasing" {
				r.safe = false
				return i, false
			} else {
				dir = "decreasing"
				if r.levels[i-1]-lvl < 1 || r.levels[i-1]-lvl > 3 {
					r.safe = false
					return i, false
				}
			}
		}

		if lvl == r.levels[i-1] {
			r.safe = false
			return i, false
		}

	}
	return 0, true
}

func main() {

	reports := []*report{}

	scanner := fileScanner()
	for scanner.Scan() {
		report := &report{}
		parts := strings.Split(scanner.Text(), " ")
		for _, val := range parts {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			report.levels = append(report.levels, int32(intVal))
		}
		breaker, safe := report.eval()
		if !safe {
			report.levels = remove(report.levels, breaker)
			report.eval()
		}
		reports = append(reports, report)
	}

	safeReports := 0
	for _, report := range reports {
		if report.safe {
			safeReports++
		}
	}

	fmt.Printf("The answer to the puzzle part 1 is: %v\n", safeReports)
	// fmt.Printf("The answer to the puzzle part 2 is: %v\n", xx)

}
