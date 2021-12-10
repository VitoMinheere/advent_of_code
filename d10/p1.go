package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// First scan outside of loop to get first depth
	scanner.Scan()

	closing_brackets := make([]string, 4)
	closing_brackets[0] = ")"
	closing_brackets[1] = "]"
	closing_brackets[2] = "}"
	closing_brackets[3] = ">"

	incomplete_lines := []string{}

	// Loop over depths
	for scanner.Scan() {
		line := scanner.Text()
		end_char := line[len(line)-1:]
		closes := false
		for _, b := range closing_brackets {
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
