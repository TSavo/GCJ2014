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

func CookieClick(C, F, X float64) float64 {
	rate := 2.0
	total_time := 0.0
	if X <= C {
		total_time = X / rate
	} else {
		for {
			if (X / rate) < ((C / rate) + (X / (rate + F))) {
				total_time += X / rate
				break
			} else {
				total_time += C / rate
				rate += F
			}
		}
	}
	return total_time
}

func main() {
	cases, _ := readLines("B-large.in")
	numCases, _ := strconv.Atoi(cases[0])
	cases = cases[1:len(cases)]
	testCases := make([]string, 0)
	fmt.Println(cases)
	for x := 0; x < numCases; x++ {
		testCases = append(testCases, cases[x])
	}
	fmt.Println(testCases)
	resultList := make([]string, 0)
	for caseNum, testCase := range testCases {
		params := strings.Split(testCase, " ")
		C, _ := strconv.ParseFloat(params[0], 32)
		F, _ := strconv.ParseFloat(params[1], 32)
		X, _ := strconv.ParseFloat(params[2], 32)
		result := fmt.Sprintf("Case #"+strconv.Itoa(caseNum+1)+": %.7f", CookieClick(C, F, X))
		fmt.Println(result)
		resultList = append(resultList, result)
	}
	writeLines(resultList, "cookieout.txt")
}
