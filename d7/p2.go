package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "sort"
    "math"
	"strconv"
	"strings"
)

// type
type MeanMedian struct {
	numbers []float64
}

// check if the total of numbers is
// odd or even
func (mm *MeanMedian) IsOdd() bool {
	if len(mm.numbers)%2 == 0 {
		return false
	}

	return true
}

func (mm *MeanMedian) CalcMean() float64 {
	total := 0.0

	for _, v := range mm.numbers {
		total += v
	}

	// IMPORTANT: return was rounded!
    // TODO fix the off by one error
	return math.Round(total / float64(len(mm.numbers))) - 1
}

func (mm *MeanMedian) CalcFuelCost(mean int, n ...float64) int {
	fuel_cost := 0
	for _, j := range mm.numbers {
		i := int(j)
		x := 0
		if i < mean {
			x = mean - i
            y := 0
            for i := 1;i <= x; i++ {
                y += i
            }
			fuel_cost += y

		} else {
			// i is larger
			y := mean - i
			x -= i
			x = -x
			y = -y
			x = y
            z := 0
            for i := 1;i <= x; i++ {
                z += i
            }
			fuel_cost += z
		}
	}
	return fuel_cost

}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	crabs := MeanMedian{}
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, ",")
		for _, v := range s {
			pos, _ := strconv.ParseFloat(v, 64)
			crabs.numbers = append(crabs.numbers, pos)

		}
	}
	mean := crabs.CalcMean()
	fmt.Print(mean)
	mean_int := int(mean)
	res := crabs.CalcFuelCost(mean_int)
	fmt.Printf("Result %s \n", strconv.Itoa(res))
}
