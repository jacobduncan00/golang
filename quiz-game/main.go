package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()
	// Pointer to a string
	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}

	r := csv.NewReader(file)
	// Parse the entire file upfront
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	problems := parseLines(lines)

	correct := 0
	// i = index p = problem
	for i, p := range problems {
		// Start at problem 1
		fmt.Printf("Problem #%d: %s = ", i+1, p.question)
		// Read in answer
		var answer string
		// Scanf will get rid of trailing spaces, etc
		// This is fine for this game
		fmt.Scanf("%s\n", &answer)
		if answer == p.answer {
			// Increment correct counter
			correct++
		}
	}

	// Potential edge cases to consider?
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
	var percentage float64 = (float64(correct) / float64(len(problems))) * 100
	fmt.Printf("Grade: %.2f%%\n", percentage)
}

func parseLines(lines [][]string) []problem {
	// Make our return value with length of number of lines in csv file as each is
	// a problem
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: strings.TrimSpace(line[0]),
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	// Status code of 1 showing something went wrong
	os.Exit(1)
}
