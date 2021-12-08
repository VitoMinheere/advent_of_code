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

type CodeMapping struct {
	code  string
	value string
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// codes = [CodeMapping{"acedgfb", 8}, CodeMapping{"cdfbe", 5}, CodeMapping{"gcdfa", 2}, CodeMapping{"fbcad", 3}, CodeMapping{"dab", 7}, CodeMapping{"cefabd", 9}, CodeMapping{"cdfgeb", 6}, CodeMapping{"eafb", 4}, CodeMapping{"cagedb", 0}, CodeMapping{"ab", 1}]
	var codes [10]CodeMapping
	codes[0] = CodeMapping{"acedgfb", "8"}
	codes[1] = CodeMapping{"cdfbe", "5"}
	codes[2] = CodeMapping{"gcdfa", "2"}
	codes[3] = CodeMapping{"fbcad", "3"}
	codes[4] = CodeMapping{"dab", "7"}
	codes[5] = CodeMapping{"cefabd", "9"}
	codes[6] = CodeMapping{"cdfgeb", "6"}
	codes[7] = CodeMapping{"eafb", "4"}
	codes[8] = CodeMapping{"cagedb", "0"}
	codes[9] = CodeMapping{"ab", "1"}

	scanner := bufio.NewScanner(file)
	// First scan outside of loop to get first depth
	simple_digits := 0
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, "| ")
		code := s[1]
		output_code := strings.Split(code, " ")
		fmt.Println(output_code)
		var output_slice []string
		for _, v := range output_code {
			value := ""
			length := len(v)
			if length == 2 {
				value = "1"
				// output_slice = append(output_slice, "1")
			}
			if length == 3 {
				value = "7"
				// output_slice = append(output_slice, "7")
			}
			if length == 4 {
				value = "4"
				// output_slice = append(output_slice, "4")
			}
			if length == 7 {
				value = "8"
				// output_slice = append(output_slice, "8")
			}
			// fmt.Printf("Length of string %s \n", strconv.Itoa(length))
			if value == "" {
				sorted_v := SortString(v)
				fmt.Printf("sorted_v: %s \n", sorted_v)
				for _, c := range codes {
					sorted_c := SortString(c.code)
					fmt.Printf("sorted_c: %s \n", sorted_c)
					if sorted_v == sorted_c {
						// fmt.Printf("Code found: %s \n", c.code)
						// fmt.Printf("Value found: %s \n", c.value)
						value = c.value
					}
				}

			}

			// fmt.Printf("Value = %s \n", value)
			output_slice = append(output_slice, value)
		}
		output := strings.Join(output_slice, "")
		fmt.Printf("Output = %s \n", output)
		output_val, _ := strconv.Atoi(output)
		simple_digits += output_val
	}
	fmt.Printf("Simple digits: %s \n", strconv.Itoa(simple_digits))

}
