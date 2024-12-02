package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

func main() {

	left := []int32{}
	right := []int32{}

	scanner := fileScanner()
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "   ")
		lInt, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		left = append(left, int32(lInt))
		rInt, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		right = append(right, int32(rInt))
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	if len(left) != len(right) {
		panic("left and right different sizes")
	}

	diffs := []int32{}
	similaritiesScores := []int32{}

	for i, l := range left {
		counter := 0
		for _, r := range right {
			if r == l {
				counter += 1
			}
		}

		if l < right[i] {
			diffs = append(diffs, right[i]-l)
		} else {
			diffs = append(diffs, l-right[i])
		}

		similaritiesScores = append(similaritiesScores, l*int32(counter))
	}

	var total int32
	for _, v := range diffs {
		total += v
	}

	fmt.Printf("The answer to the puzzle part 1 is: %v\n", total)

	var similaritiesScoreTotal int32

	for _, v := range similaritiesScores {
		similaritiesScoreTotal += v
	}

	fmt.Printf("The answer to the puzzle part 2 is: %v\n", similaritiesScoreTotal)

}
