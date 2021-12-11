package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	row int
	col int
}

type Octopus struct {
	energy     int
	flashed    bool
	neighbours []Pos
}

type Octi struct {
	specimens [][]Octopus
	flashes   int
}

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func getOctiMap(rawBytes []byte) Octi {
	octi := Octi{}
	octi.specimens = make([][]Octopus, 10, 10)
	octi.flashes = 0

	lines := strings.Split(string(rawBytes), "\n")
	for i, line := range lines {
		s := strings.Split(line, "")
		row := make([]Octopus, len(s))
		for j, v := range s {
			number, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			row[j].energy = number
			row[j].flashed = false
			row[j].neighbours = findNeighbours(octi, i, j)
		}
		octi.specimens[i] = row
	}
	return octi
}

// All octi get a +1
func stepOne(octi Octi) Octi {
	for i, r := range octi.specimens {
		for j := range r {
			octi.specimens[i][j].energy++
		}
	}
	return octi
}

// All octi get a +1
func stepTwo(octi Octi) Octi {
	for i, r := range octi.specimens {
		for j, v := range r {
			if v.energy == 9 && v.flashed == false {
				octi = flashOctopus(octi, i, j)
			}
		}
	}
	return octi
}

func stepThree(octi Octi) Octi {
	for i, r := range octi.specimens {
		for j := range r {
			octi.specimens[i][j].flashed = false
		}
	}
	return octi
}

func findNeighbours(array Octi, r int, c int) []Pos {
	neighbours := []Pos{}
	// fmt.Printf("Cur pos = %s, %s \n", strconv.Itoa(r), strconv.Itoa(c))

	row_limit := 9 //len(array.specimens) - 1
	if row_limit > 0 {
		column_limit := 9 //len(array.specimens[0]) - 1
		for x := Max(0, r-1); x <= Min(r+1, row_limit); x++ {
			// if x < 2 {
			// 	fmt.Printf("x = %s \n", strconv.Itoa(x))
			// }
			for y := Max(0, c-1); y <= Min(c+1, column_limit); y++ {
				// if y < 2 {
				// 	fmt.Printf("y = %s \n", strconv.Itoa(y))
				// }
				// fmt.Printf(" x = %s and y = %s \n", strconv.Itoa(x), strconv.Itoa(y))
				if x != r || y != c {
					neighbours = append(neighbours, Pos{x, y})
				}
			}
		}
	}
	// fmt.Println(neighbours)
	return neighbours
}

func flashOctopus(octi Octi, row int, col int) Octi {
	cur_octopus := &octi.specimens[row][col]
	cur_octopus.energy = 0
	cur_octopus.flashed = true
	octi.flashes++

	for _, n := range cur_octopus.neighbours {
		octi.specimens[n.row][n.col].energy++
	}

	return octi
}

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rawBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	octi := getOctiMap(rawBytes)
	// fmt.Println(octi.specimens[0][3])

	for n := 0; n < 2; n++ {
		fmt.Printf("Step %s \n", strconv.Itoa(n))
		octi = stepOne(octi)
		for _, r := range octi.specimens {
			fmt.Println(r)
			fmt.Printf("\n")
			//TODO Check with example output

		}
		octi = stepTwo(octi)
		octi = stepThree(octi)

	}

	fmt.Println(octi.flashes)
}
