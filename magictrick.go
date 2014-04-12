// +build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

type Cards [][]int

func MakeCards(cards []string) *Cards {
	deck := make(Cards, 4)
	for x := 0; x < 4; x++ {
		deck[x] = make([]int, 4)
	}
	for rowNum, row := range cards {
		for x, card := range strings.Split(row, " ") {
			deck[rowNum][x], _ = strconv.Atoi(card)
		}
	}
	return &deck
}

func main() {

	cases, _ := readLines("A-small-attempt0.in")
	numCases, _ := strconv.Atoi(cases[0])
	cases = cases[1:len(cases)]
	testCases := make([][]string, 0)
	fmt.Println(cases)
	for x := 0; x < numCases; x++ {
		testCases = append(testCases, cases[x*10:(x*10)+10])
	}
	fmt.Println(testCases)
	resultList := make([]string, 0)
	for caseNum, hand := range testCases {
		firstGuess, _ := strconv.Atoi(hand[0])
		secondGuess, _ := strconv.Atoi(hand[5])
		deck1 := MakeCards(hand[1:5])
		deck2 := MakeCards(hand[6:10])
		found := false
		badMagician := false
		magic := 0
		for _, first := range (*deck1)[firstGuess-1] {
			for _, second := range (*deck2)[secondGuess-1] {
				if first == second {
					fmt.Println(first)
					if found {
						badMagician = true
					} else {
						magic = first
						found = true
					}
				}
			}
		}
		result := "Case #" + strconv.Itoa(caseNum +1) + ": "
		if badMagician {
			result += "Bad magician!"
		} else if found {
			result += strconv.Itoa(magic)
		} else {
			result += "Volunteer cheated!"
		}
		fmt.Println(firstGuess, secondGuess, deck1, deck2)
		fmt.Println(result)
		resultList = append(resultList, result)
	}
	writeLines(resultList, "out.txt")
}
