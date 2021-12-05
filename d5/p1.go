package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x        int
	y        int
	has_vent bool
	vents    int
}

// Help-function to insert at index
func insert(a []Coord, index int, value Coord) []Coord {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func createVents(coords []Coord, distance int, horizontal bool) []Coord {
	start := coords[0]
	end := coords[1]

	fmt.Printf("Distance = %s \n", strconv.Itoa(distance))

	if horizontal == true {
		if distance < 0 {
			fmt.Print("Negative vents \n")
			for i := 1; i < -distance; i++ {
				vent := Coord{x: end.x - i, y: end.y, has_vent: true, vents: 1}
				coords = insert(coords, i, vent)
			}

		} else {
			for i := 1; i < distance; i++ {
				vent := Coord{x: start.x - i, y: start.y, has_vent: true, vents: 1}
				coords = insert(coords, i, vent)
			}
		}

	} else {
		if distance < 0 {
			fmt.Print("Negative vents \n")
			for i := 1; i < -distance; i++ {
				vent := Coord{x: end.x, y: end.y - i, has_vent: true, vents: 1}
				coords = insert(coords, i, vent)
			}
		} else {
			for i := 1; i < distance; i++ {
				vent := Coord{x: start.x, y: start.y - i, has_vent: true, vents: 1}
				coords = insert(coords, i, vent)
			}
		}

	}
	return coords
}

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create board to build upon
	// board := make([][]int, 9, 9)

	scanner := bufio.NewScanner(file)

	// vent_coords := []Coord{}

	// Loop over depths
	for scanner.Scan() {
		line := scanner.Text()
		// Strip out the arrow and split coords
		s := strings.Split(line, " -> ")
		line_coords := []Coord{}
		// Loop over start and end
		for _, v := range s {
			// fmt.Println(v)
			// Get x and y values
			coords := strings.Split(v, ",")
			x, err := strconv.Atoi(coords[0])
			y, err := strconv.Atoi(coords[1])
			if err != nil {
				log.Fatal(err)
			}
			coord := Coord{x: x, y: y, has_vent: true, vents: 1}
			line_coords = append(line_coords, coord)
		}
		// fmt.Println(line_coords)
		// For now, only consider horizontal and vertical lines: lines where either x1 = x2 or y1 = y2.
		if line_coords[0].x == line_coords[1].x {
			fmt.Println(s)
			vert_dis := line_coords[0].y - line_coords[1].y
			fmt.Printf("Vertical vents %s \n", strconv.Itoa(vert_dis))
			vents := createVents(line_coords, vert_dis, false)
			fmt.Println(vents)
		}
		if line_coords[0].y == line_coords[1].y {
			fmt.Println(s)
			hor_dis := line_coords[0].x - line_coords[1].x
			fmt.Printf("Horizontal vents %s \n", strconv.Itoa(hor_dis))
			vents := createVents(line_coords, hor_dis, true)
			fmt.Println(vents)

		}

	}
}
