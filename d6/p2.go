package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
    "runtime"
	"strconv"
	"strings"
)

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number of garage collection cycles completed.
func PrintMemUsage() {
        var m runtime.MemStats
        runtime.ReadMemStats(&m)
        // For info on each, see: https://golang.org/pkg/runtime/#MemStats
        fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
        fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
        fmt.Printf("\tMAlloc = %v MiB", bToMb(m.Mallocs))
        fmt.Printf("\tHeapAlloc = %v MiB", bToMb(m.HeapAlloc))
        fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
        fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// First scan outside of loop to get first depth
	lantern_fish := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		// Strip out the arrow and split coords
		s := strings.Split(line, ",")
		for _, v := range s {
			days, _ := strconv.Atoi(v)
			lantern_fish = append(lantern_fish, days)
		}
	}
	fmt.Println(lantern_fish)

	// Simulate days
	for i := 0; i <= 256; i++ {
		fmt.Printf("Day %s \n", strconv.Itoa(i))
		fmt.Printf("Amount of fish start of day: %s \n", strconv.Itoa(len(lantern_fish)))
		// fmt.Println(lantern_fish)
		for j := range lantern_fish {
			if lantern_fish[j] == 0 {
				lantern_fish = append(lantern_fish, 8)
				lantern_fish[j] = 6
			} else {
				lantern_fish[j]--
			}
		}
		// if len(new_fishes) > 0 {
		// 	lantern_fish = append(lantern_fish, new_fishes...)
		// }
        runtime.GC()
        PrintMemUsage()
		// fmt.Printf("Amount of fish after day: %s \n", strconv.Itoa(len(lantern_fish)))
	}
}
