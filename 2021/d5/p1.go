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

func createField(x int, y int) [][]Coord {
	field := make([][]Coord, x, y)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			field[i] = append(field[i], Coord{x: i, y: j, has_vent: false, vents: 0})
		}
	}
	return field
}

func createVents(coords []Coord, distance int, horizontal bool) []Coord {
	start := coords[0]
	end := coords[1]

	// fmt.Printf("Distance = %s \n", strconv.Itoa(distance))

	if horizontal == true {
		if distance < 0 {
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

func placeVents(field [][]Coord, vents []Coord) ([][]Coord, int) {
	lines_overlap := 0
	for _, v := range vents {
		// Store current coord values
		sel_field := field[v.x][v.y]
		field[v.x][v.y] = v
		if sel_field.vents > 0 {
			field[v.x][v.y].vents = sel_field.vents + 1
		}
		if field[v.x][v.y].vents == 2 {
			lines_overlap++
		}
		// fmt.Println(field[v.x][v.y])
	}
	return field, lines_overlap
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create board to build upon
	field := createField(1000, 1000)

	scanner := bufio.NewScanner(file)

	all_vent_coords := []Coord{}

	// Loop over depths
	for scanner.Scan() {
		line := scanner.Text()
		// Strip out the arrow and split coords
		s := strings.Split(line, " -> ")
		line_coords := []Coord{}
		vents := []Coord{}
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
		// For now, only consider horizontal and vertical lines: lines where either x1 = x2 or y1 = y2.
		if line_coords[0].x == line_coords[1].x {
			vert_dis := line_coords[0].y - line_coords[1].y
			// fmt.Printf("Vertical vents %s \n", strconv.Itoa(vert_dis))
			vents = createVents(line_coords, vert_dis, false)
		}
		if line_coords[0].y == line_coords[1].y {
			hor_dis := line_coords[0].x - line_coords[1].x
			// fmt.Printf("Horizontal vents %s \n", strconv.Itoa(hor_dis))
			vents = createVents(line_coords, hor_dis, true)
		}
		all_vent_coords = append(all_vent_coords, vents...)

	}
	_, overlap := placeVents(field, all_vent_coords)

	// for _, v := range field_with_vents {
	// 	fmt.Println(v)
	// }
	fmt.Println(overlap)
}
