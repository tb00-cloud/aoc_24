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

	// b = []byte("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))")

	re := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	cumulative := 0

	blocks := strings.Split(string(b), "do()")

	for _, block := range blocks {
		in := strings.Index(block, `don't()`)
		nb := block
		if in != -1 {
			nb = block[0:in]
		}

		for _, eq := range re.FindAllString(nb, -1) {
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
	}

	fmt.Printf("The answer is: %v\n", cumulative)

}
