package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var depth int
	var horizontal int
	var aim int

	scanner := bufio.NewScanner(file)

	// Loop over depths
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "forward") {
			words := strings.Fields(line)
			hor_value := words[len(words)-1]
			hor_int, _ := strconv.Atoi(hor_value)
			horizontal += hor_int
			depth += (hor_int * aim)
		}

		if strings.Contains(line, "up") {
			words := strings.Fields(line)
			up_value := words[len(words)-1]
			up_int, _ := strconv.Atoi(up_value)
			aim -= up_int
		}

		if strings.Contains(line, "down") {
			words := strings.Fields(line)
			down_value := words[len(words)-1]
			down_int, _ := strconv.Atoi(down_value)
			aim += down_int
		}
	}
	fmt.Println(horizontal * depth)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
