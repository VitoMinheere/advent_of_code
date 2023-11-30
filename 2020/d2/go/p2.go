package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	defer timeTrack(time.Now(), "p2")
	data := parse()

	valid_passwords := 0

	for _, v := range data {
		pos1, pos2, _ := strings.Cut(v, "-")
		pos2, rest, _ := strings.Cut(pos2, " ")
		char, password, _ := strings.Cut(rest, ": ")

		pos_1, _ := strconv.Atoi(pos1)
		pos_2, _ := strconv.Atoi(pos2)

		pos1_char := string(password[pos_1-1])
		pos2_char := string(password[pos_2-1])
		if (pos1_char == char || pos2_char == char) && pos1_char != pos2_char {
			valid_passwords++
		}
		// if pos1_char == char {
		// 	first_correct := true
		// 	if pos2_char == char && !first_correct {
		// 		valid_passwords++
		// 	}
		// 	if pos2_char != char && first_correct {
		// 		valid_passwords++
		// 	}
		// }
	}
	fmt.Println(valid_passwords)
	return

}
