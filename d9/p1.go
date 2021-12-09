package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			count += 1
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func getHeightMap(rawBytes []byte) [][]int {
	height_map := make([][]int, 100, 100)

	lines := strings.Split(string(rawBytes), "\n")
	for i, line := range lines {
		s := strings.Split(line, "")
		row := make([]int, len(s))
		for j, v := range s {
			number, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			row[j] = number
		}
		height_map[i] = row
		// fmt.Println(row)
	}
	return height_map
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rawBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	height_map := getHeightMap(rawBytes)

	lowest_points := []int{}

	for i, r := range height_map {
		// Set to highest value
		left := 9
		right := 9
		up := 9
		down := 9

		// Loop over numbers in row
		for j, v := range r {
			if v == 0 {
				fmt.Printf("Found low point %s at row %s col %s \n", strconv.Itoa(v), strconv.Itoa(i), strconv.Itoa(j))
				lowest_points = append(lowest_points, v)
				continue
			}
			if j != 0 {
				left = r[j-1]
			}
			if j != len(r)-1 {
				right = r[j+1]
			}

			// top or bottom row
			if i != 0 {
				up = height_map[i-1][j]

			}
			if i < len(height_map)-1 {
				down = height_map[i+1][j]
			}

			if v < left && v < right && v < up && v < down {
				fmt.Printf("Found low point %s at row %s col %s \n", strconv.Itoa(v), strconv.Itoa(i), strconv.Itoa(j))
				lowest_points = append(lowest_points, v)
			}
		}
	}
	fmt.Println(lowest_points)

	result := 0
	for _, v := range lowest_points {
		result += (v + 1)
	}
	fmt.Printf("Result %s \n", strconv.Itoa(result))

}
