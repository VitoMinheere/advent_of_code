package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

func parse() []string {
	scanner := bufio.NewScanner(os.Stdin)

	list := []string{}
	for scanner.Scan() {
		line := scanner.Text()

		list = append(list, line)
	}
	return list
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
