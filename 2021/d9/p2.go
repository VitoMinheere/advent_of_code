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

type BasinTile struct {
	value int
	row   int
	col   int
}

type Basin struct {
	tiles []BasinTile
	value int
}

func contains(s []BasinTile, e BasinTile) bool {
	for _, a := range s {
		if a.row == e.row && a.col == e.col {
			return true
		}
	}
	return false
}

func tileAlreadyInBasin(basins []Basin, cur_tile BasinTile) bool {
	result := false
	for _, b := range basins {
		if contains(b.tiles, cur_tile) {
			result = true
			fmt.Printf("%s already in found basins \n", cur_tile)
			break
		} else {
			break
		}
	}
	return result
}

func checkDirection(direction string, row int, col int, height_map [][]int, basins []Basin, basin Basin) Basin {
	value := 0
	offset := 1
	neighbour := 9
	tile := BasinTile{}
	for value != 9 {
		if direction == "r" {
			if col+offset < len(height_map[row]) {
				neighbour = height_map[row][col+offset]
				tile = BasinTile{neighbour, row, col + offset}
			} else {
				break
			}
		}
		if direction == "l" {
			if col-offset > 0 {
				neighbour = height_map[row][col-offset]
				tile = BasinTile{neighbour, row, col - offset}
			} else {
				break
			}
		}
		if direction == "d" {
			if row+offset < len(height_map) {
				neighbour = height_map[row+offset][col]
				tile = BasinTile{neighbour, row + offset, col}
			} else {
				break
			}
		}
		if neighbour != 9 {
			fmt.Printf("Found %s \n", neighbour)
			cur_basin := []Basin{}
			cur_basin = append(cur_basin, basin)

			if !tileAlreadyInBasin(basins, tile) && !tileAlreadyInBasin(cur_basin, tile) {
				fmt.Printf("Adding %s \n", tile)
				basin.tiles = append(basin.tiles, tile)
				offset++
			} else {
				break
			}
		} else {
			break
		}
	}
	return basin
}

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
	height_map := make([][]int, 5, 5)

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
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rawBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	height_map := getHeightMap(rawBytes)

	basins := []Basin{}

	for i, r := range height_map {
		fmt.Printf(" i = %s \n", strconv.Itoa(i))
		// Set to highest value
		// left := 9
		// up := 9
		pos_9 := 0
		fmt.Print(pos_9)

		// Loop over each numbers in matrix
		for j, v := range r {
			// right := 0
			fmt.Print("New basin \n")
			fmt.Printf("Checking value %s row %s col %s \n", strconv.Itoa(v), strconv.Itoa(i), strconv.Itoa(j))
			// Append the first value
			cur_tile := BasinTile{}
			if v != 9 {
				cur_tile = BasinTile{v, i, j}
			} else {
				continue
			}

			basin := Basin{}

			// Check if number is not yet in any basin
			if len(basins) > 0 {
				if tileAlreadyInBasin(basins, cur_tile) {
					fmt.Print("Tile already in basins")
					continue
				}
				// else {
				// 	basin.tiles = append(basin.tiles, cur_tile)
				// }
			} else {
				basin.tiles = append(basin.tiles, cur_tile)
			}

			fmt.Print("start checking right \n")
			basin = checkDirection("r", i, j, height_map, basins, basin)

			// For each horizontal value go down as well
			fmt.Print("start checking down \n")
			for _, t := range basin.tiles {
				basin = checkDirection("d", t.row, t.col, height_map, basins, basin)
			}

			fmt.Print("start checking left \n")
			for _, t := range basin.tiles {
				basin = checkDirection("l", t.row, t.col, height_map, basins, basin)
			}

			if len(basin.tiles) > 1 {
				for _ = range basin.tiles {
					basin.value += 1
				}
				fmt.Printf("basins = %s \n", strconv.Itoa(len(basins)))
				fmt.Println(basin)
				basins = append(basins, basin)
			}
		}
	}

	// result := 0
	// for _, v := range lowest_points {
	// 	result += (v + 1)
	// }
	// fmt.Printf("Result %s \n", strconv.Itoa(result))

}
