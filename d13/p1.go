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

func getHighestNumbers(rawBytes []byte) (int, int) {
	max_x := 0
	max_y := 0

	lines := strings.Split(string(rawBytes), "\n")
	for _, line := range lines {
		if len(line) > 0 && len(line) <= 9 {
			s := strings.Split(line, ",")
			x, _ := strconv.Atoi(s[0])
			y, _ := strconv.Atoi(s[1])

			if x > max_x {
				max_x = x
			}

			if y > max_y {
				max_y = y
			}
		}
	}
	return max_x, max_y
}

func applyDots(rawBytes []byte) [][]int {
	paper := make([][]int, 895)
	for i := 0; i < 895; i++ {
		row := make([]int, 1311)
		paper[i] = row
	}

	lines := strings.Split(string(rawBytes), "\n")
	for _, line := range lines {
		if len(line) > 0 && len(line) <= 9 {
			s := strings.Split(line, ",")
			if len(s) <= 9 && len(s) > 0 {
				fmt.Println(s)
				x, _ := strconv.Atoi(s[0])
				y, _ := strconv.Atoi(s[1])
				paper[y][x] = 1
			} else {
				continue
			}

		}
	}
	return paper
}

func foldPaper(paper [][]int, fold int) [][]int {
	first_half := make([][]int, fold+1)
	first_half = paper[:fold]
	// fmt.Print(len(first_half))

	second_half := make([][]int, fold)
	second_half = paper[fold+1:]
	// fmt.Print(len(second_half))

	// Only vertical folds first
	for i, r := range second_half {
		y := -(i - fold)
		fmt.Printf("Row %s is folded y %s \n", strconv.Itoa(i), strconv.Itoa(y))
		for j, v := range r {
			if v == 1 {
				first_half[y-1][j] = v
			}
		}
	}

	return first_half
}

func countDots(paper [][]int) int {
	sum_dots := 0
	for _, r := range paper {
		for _, v := range r {
			sum_dots += v
		}
	}
	return sum_dots
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

	paper := applyDots(rawBytes)

	paper = foldPaper(paper, 655)
	// for _, r := range paper {
	// 	fmt.Println(r)
	// }

	result := countDots(paper)
	fmt.Printf("Result = %s \n", strconv.Itoa(result))
}
