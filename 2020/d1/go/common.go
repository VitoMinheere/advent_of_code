package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"
)

func parse() []int {
	scanner := bufio.NewScanner(os.Stdin)

	list := []int{}
	for scanner.Scan() {
		line := scanner.Text()

		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		list = append(list, num)
	}
	return list
}

func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    log.Printf("%s took %s", name, elapsed)
}