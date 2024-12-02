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
	levels                []int32
	outlierLevelPositions []int
	safe                  bool
}

func (r *report) eval(faultTolerance int) {

	dir := ""

	r.safe = true
	for i, lvl := range r.levels {
		if i == 0 {
			continue
		}

		// If current bigger than last
		if lvl > r.levels[i-1] {
			if dir == "decreasing" {
				r.outlierLevelPositions = append(r.outlierLevelPositions, i)
				continue
			} else {
				dir = "increasing"
				if lvl-r.levels[i-1] < 1 || lvl-r.levels[i-1] > 3 {
					r.outlierLevelPositions = append(r.outlierLevelPositions, i)
					continue
				}
			}
		}

		if lvl < r.levels[i-1] {
			if dir == "increasing" {
				r.outlierLevelPositions = append(r.outlierLevelPositions, i)
				continue
			} else {
				dir = "decreasing"
				if r.levels[i-1]-lvl < 1 || r.levels[i-1]-lvl > 3 {
					r.outlierLevelPositions = append(r.outlierLevelPositions, i)
					continue
				}
			}
		}

		if lvl == r.levels[i-1] {
			r.outlierLevelPositions = append(r.outlierLevelPositions, i)
			continue
		}
	}

	fmt.Printf("report outlier count: %v\n", r.outlierLevelPositions)

	if len(r.outlierLevelPositions) > faultTolerance {
		r.safe = false
	}

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
		reports = append(reports, report)
	}

	p1SafeReports := 0
	for _, report := range reports {
		report.eval(0)
		if report.safe {
			p1SafeReports++
		}
	}
	fmt.Printf("The answer to the puzzle part 1 is: %v\n", p1SafeReports)

	p2SafeReports := 0
	for _, report := range reports {
		report.outlierLevelPositions = []int{}
		report.eval(1)
		if report.safe {
			p2SafeReports++
		}
	}
	fmt.Printf("The answer to the puzzle part 2 is: %v\n", p2SafeReports)

}
