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

type Fold struct {
    axis    string
    index   int
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
    // fmt.Printf("highest x : %s \n", strconv.Itoa(max_x))
    // fmt.Printf("highest y : %s \n", strconv.Itoa(max_y))
	return max_x, max_y
}

func getFolds(rawBytes []byte) []Fold {
    folds := []Fold{}

	lines := strings.Split(string(rawBytes), "\n")
	for _, line := range lines {
		if strings.Contains(line, "fold along") {
			s := strings.Split(line, "=")
            axis := s[0][len(s[0])-1:]
            index, _ := strconv.Atoi(s[1])
            folds = append(folds, Fold{axis, index})
		}
	}
	return folds
}

func applyDots(rawBytes []byte, cols int, rows int) [][]int {
	paper := make([][]int, rows+1)
	for i := 0; i <= rows; i++ {
		row := make([]int, cols+1)
		paper[i] = row
	}

	lines := strings.Split(string(rawBytes), "\n")
	for _, line := range lines {
		if len(line) > 0 && len(line) <= 9 {
			s := strings.Split(line, ",")
			if len(s) <= 9 && len(s) > 0 {
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

func foldPaper(paper [][]int, fold int, fold_type string) [][]int {
    folded := [][]int{}

    if fold_type == "y" {
        first_half := make([][]int, fold+1)
        first_half = paper[:fold]

        second_half := make([][]int, fold)
        second_half = paper[fold+1:]

        for i, r := range second_half {
            y := -(i - fold)
            // fmt.Printf("Row %s is folded y %s \n", strconv.Itoa(i), strconv.Itoa(y))
            for j, v := range r {
                if v == 1 {
                    first_half[y-1][j] = v
                }
            }
        }
        folded = first_half
    }
    if fold_type == "x" {
        first_half := make([][]int, len(paper)+1)
        for i, r := range paper {
            first_half[i] = r[:fold]
        }

        second_half := make([][]int, len(paper)+1)
        for i, r := range paper {
            second_half[i] = r[fold+1:]
        }

        for i, r := range second_half {
            for j, v := range r {
                x := -(j - fold)
                // fmt.Printf("Col %s is folded x %s \n", strconv.Itoa(j), strconv.Itoa(x))
                if v == 1 {
                    first_half[i][x-1] = v
                }
            }
        }
        folded = first_half
    }

	return folded
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

    x, y := getHighestNumbers(rawBytes)
    folds := getFolds(rawBytes)
	paper := applyDots(rawBytes, x, y)

	for _, r := range folds {
	    paper = foldPaper(paper, r.index, r.axis)
	}

	result := countDots(paper)
	fmt.Printf("Result p1 = %s \n", strconv.Itoa(result))

    // result_p2
    for _, r := range paper {
        text := ""
        for _, v := range r {
            if v == 0 {
                text += " "
            }
            if v == 1 {
                text += "#"
            }
        }
        fmt.Println(text)
    }
}
