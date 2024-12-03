package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	b, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	cumulative := 0
	for _, eq := range re.FindAllString(string(b), -1) {
		eq = strings.Replace(eq, "mul(", "", -1)
		eq = strings.Replace(eq, ")", "", -1)
		parts := strings.Split(eq, ",")
		if len(parts) != 2 {
			panic("parts not 2")
		}
		l, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		cumulative += l * r
	}

	fmt.Printf("The answer is: %v\n", cumulative)

}
