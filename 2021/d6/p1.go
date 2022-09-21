package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Fish struct {
	days_left int
}

func createFish(days_left int) Fish {
	new_fish := Fish{days_left: days_left}
	return new_fish
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// First scan outside of loop to get first depth
	lantern_fish := []Fish{}
	for scanner.Scan() {
		line := scanner.Text()
		// Strip out the arrow and split coords
		s := strings.Split(line, ",")
		for _, v := range s {
			days, _ := strconv.Atoi(v)
			fish := Fish{days} //createFish(days)
			lantern_fish = append(lantern_fish, fish)

		}
	}
	fmt.Println(lantern_fish)

	// Simulate days
	for i := 0; i <= 256; i++ {
		// fmt.Printf("Day %s \n", strconv.Itoa(i))
		// fmt.Printf("Amount of fish start of day: %s \n", strconv.Itoa(len(lantern_fish)))
		// fmt.Println(lantern_fish)
		new_fishes := []Fish{}
		for j := range lantern_fish {
			if lantern_fish[j].days_left == 0 {
				new_fishes = append(new_fishes, Fish{8})
				lantern_fish[j].days_left = 6
			} else {
				lantern_fish[j].days_left--
			}
		}
		if len(new_fishes) > 0 {
			lantern_fish = append(lantern_fish, new_fishes...)
		}
		// fmt.Printf("Amount of fish after day: %s \n", strconv.Itoa(len(lantern_fish)))
	}
	fmt.Printf("Amount of fish start of day: %s \n", strconv.Itoa(len(lantern_fish)))
}
