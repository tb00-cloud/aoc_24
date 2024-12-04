package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func fileScanner() *bufio.Scanner {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewScanner(file)
}

type grid struct {
	rows [][]string
}

func (g *grid) fromFile() {
	scanner := fileScanner()
	for scanner.Scan() {
		row := []string{}
		for _, s := range scanner.Text() {
			row = append(row, string(s))
		}
		g.rows = append(g.rows, row)
	}
}

type coords struct {
	row  int
	col  int
	hDir string
	vDir string
}

func (g *grid) try(row int, col int, rowOffset int, colOffset int, char string) bool {
	if row+rowOffset < 0 || row+rowOffset > len(g.rows[row])-1 {
		return false
	}
	if col+colOffset < 0 || col+colOffset > len(g.rows)-1 {
		return false
	}

	return g.rows[row+rowOffset][col+colOffset] == char
}

func (g *grid) walk() int {

	hits := 0

	for rowI, row := range g.rows {
		for colI := range row {
			if g.try(rowI, colI, 0, 0, "X") {
				// Up
				if g.try(rowI, colI, -1, 0, "M") {
					if g.try(rowI, colI, -2, 0, "A") {
						if g.try(rowI, colI, -3, 0, "S") {
							hits++
						}
					}
				}
				// down
				if g.try(rowI, colI, +1, 0, "M") {
					if g.try(rowI, colI, +2, 0, "A") {
						if g.try(rowI, colI, +3, 0, "S") {
							hits++
						}
					}
				}
				// left
				if g.try(rowI, colI, 0, -1, "M") {
					if g.try(rowI, colI, 0, -2, "A") {
						if g.try(rowI, colI, 0, -3, "S") {
							hits++
						}
					}
				}
				// right
				if g.try(rowI, colI, 0, +1, "M") {
					if g.try(rowI, colI, 0, +2, "A") {
						if g.try(rowI, colI, 0, +3, "S") {
							hits++
						}
					}
				}
				// Up-left
				if g.try(rowI, colI, -1, -1, "M") {
					if g.try(rowI, colI, -2, -2, "A") {
						if g.try(rowI, colI, -3, -3, "S") {
							hits++
						}
					}
				}
				// Up-right
				if g.try(rowI, colI, -1, +1, "M") {
					if g.try(rowI, colI, -2, +2, "A") {
						if g.try(rowI, colI, -3, +3, "S") {
							hits++
						}
					}
				}
				// Down-left
				if g.try(rowI, colI, +1, -1, "M") {
					if g.try(rowI, colI, +2, -2, "A") {
						if g.try(rowI, colI, +3, -3, "S") {
							hits++
						}
					}
				}
				// Down-right
				if g.try(rowI, colI, +1, +1, "M") {
					if g.try(rowI, colI, +2, +2, "A") {
						if g.try(rowI, colI, +3, +3, "S") {
							hits++
						}
					}
				}
			}
		}
	}
	return hits
}

func main() {

	grd := &grid{}
	grd.fromFile()
	fmt.Printf("The answer is: %v\n", grd.walk())

}
