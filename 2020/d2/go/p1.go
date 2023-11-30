package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	defer timeTrack(time.Now(), "p1")
	data := parse()

	valid_passwords := 0

	for _, v := range data {
		min_s, max_s, _ := strings.Cut(v, "-")
		max_s, rest, _ := strings.Cut(max_s, " ")
		char, password, _ := strings.Cut(rest, ": ")

		min, _ := strconv.Atoi(min_s)
		max, _ := strconv.Atoi(max_s)
		amount := strings.Count(password, char)

		if amount >= min && amount <= max {
			valid_passwords++
		}
	}
	fmt.Println(valid_passwords)
	return

}
