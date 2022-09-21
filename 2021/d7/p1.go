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

func (mm *MeanMedian) CalcMedian(n ...float64) float64 {
	sort.Float64s(mm.numbers) // sort the numbers

	mNumber := len(mm.numbers) / 2

	if mm.IsOdd() {
		return mm.numbers[mNumber]
	}

	return (mm.numbers[mNumber-1] + mm.numbers[mNumber]) / 2
}

func (mm *MeanMedian) CalcFuelCost(median int, n ...float64) int {
	fuel_cost := 0
	for _, j := range mm.numbers {
		i := int(j)
		fmt.Printf("number %s \n", strconv.Itoa(int(i)))
		x := 0
		if i < median {
			x = median - i
			// fmt.Printf("Fuel cost += %s \n", strconv.Itoa(x))
			fuel_cost += x

		} else {
			// i is larger
			y := median - i
			x -= i
			x = -x
			y = -y
			x = y
			// fmt.Printf("Fuel cost += %s \n", strconv.Itoa(x))
			fuel_cost += x
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
	// First scan outside of loop to get first depth
	crabs := MeanMedian{}
	for scanner.Scan() {
		line := scanner.Text()
		// Strip out the arrow and split coords
		s := strings.Split(line, ",")
		for _, v := range s {
			pos, _ := strconv.ParseFloat(v, 64)
			crabs.numbers = append(crabs.numbers, pos)

		}
	}
	median := crabs.CalcMedian()
	fmt.Print(median)
	median_int := int(median)
	res := crabs.CalcFuelCost(median_int)
	fmt.Printf("Result %s \n", strconv.Itoa(res))
}
