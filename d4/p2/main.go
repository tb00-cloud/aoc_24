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
	row int
	col int
}

func (g *grid) checkNear(char string, row int, col int) []coords {
	rowMax := len(g.rows) - 1
	colMax := len(g.rows[row]) - 1

	result := []coords{}

	if row > 0 {
		// UP left
		if col > 0 {
			if g.rows[row-1][col-1] == char {
				result = append(result, coords{
					row: row - 1,
					col: col - 1,
				})
			}
		}
		// UP right
		if col < colMax {
			if g.rows[row-1][col+1] == char {
				result = append(result, coords{
					row: row - 1,
					col: col + 1,
				})
			}
		}
	}

	if row < rowMax {
		//DOWN left
		if col > 0 {
			if g.rows[row+1][col-1] == char {
				result = append(result, coords{
					row: row + 1,
					col: col - 1,
				})
			}
		}
		//DOWN right
		if col < colMax {
			if g.rows[row+1][col+1] == char {
				result = append(result, coords{
					row: row + 1,
					col: col + 1,
				})
			}
		}
	}

	return result
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
		for colI, char := range row {
			if char == "A" {
				xHits := 0
				opts := g.checkNear("M", rowI, colI)
				if len(opts) != 2 {
					continue
				} else {
					for _, opt := range opts {
						if opt.row > rowI && opt.col > colI {
							if g.try(rowI, colI, -1, -1, "S") {
								xHits++
							}
						}
						if opt.row > rowI && opt.col < colI {
							if g.try(rowI, colI, -1, +1, "S") {
								xHits++
							}

						}
						if opt.row < rowI && opt.col > colI {
							if g.try(rowI, colI, +1, -1, "S") {
								xHits++
							}
						}
						if opt.row < rowI && opt.col < colI {
							if g.try(rowI, colI, +1, +1, "S") {
								xHits++
							}
						}
					}
					if xHits == 2 {
						hits++
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
