package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Bracket struct {
	open   string
	close  string
	points int
}

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// First scan outside of loop to get first depth
	scanner.Scan()

	brackets := make([]Bracket, 4)
	brackets[0] = Bracket{"(", ")", 3}
	brackets[1] = Bracket{"[", "]", 57}
	brackets[2] = Bracket{"{", "}", 1197}
	brackets[3] = Bracket{"<", ">", 25137}

	incomplete_lines := []string{}

	// Loop over depths
	for scanner.Scan() {
		line := scanner.Text()
		end_char := line[len(line)-1:]
		closes := false
		for _, b := range brackets {
			if end_char == b {
				closes = true
			}
		}
		if closes == true {
			fmt.Println(line)
			incomplete_lines = append(incomplete_lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
