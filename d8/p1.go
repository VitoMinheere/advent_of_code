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

	scanner := bufio.NewScanner(file)
	// First scan outside of loop to get first depth
	simple_digits := 0
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, "| ")
		code := s[1]
		output_code := strings.Split(code, " ")
		fmt.Println(output_code)
		for _, v := range output_code {
			length := len(v)
			// fmt.Printf("Length of string %s \n", strconv.Itoa(length))

			if length == 4 || length == 2 || length == 3 || length == 7 {
				simple_digits++
			}
		}
	}
	fmt.Printf("Simple digits: %s \n", strconv.Itoa(simple_digits))

}
