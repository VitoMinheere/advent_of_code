package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Field struct {
	number string
	marked bool
}

type Row struct {
	fields []Field
}

type Board struct {
	rows []Row
}

type Game struct {
	boards []Board
	winner Board
}

func createGame(rawBytes []byte) Game {
	// var numbers []string

	lines := strings.Split(string(rawBytes), "\n")
	game := Game{}
	board := Board{}
	row_num := 0
	for i, line := range lines {
		if line == "" {
			continue
		}
		// Boards start at line 3
		// Boards are 5 lines long with a line break in between
		if i >= 2 && row_num <= 4 {
			fmt.Println(line)
			space := regexp.MustCompile(`\s+`)
			// Remove double space for single char numbers
			stripped := space.ReplaceAllString(line, " ")
			s := strings.Fields(stripped)
			row := Row{}
			for _, v := range s {
				num := Field{v, false}
				row.fields = append(row.fields, num)
			}
			board.rows = append(board.rows, row)
			row_num++
			// fmt.Printf("row num: %s \n", strconv.Itoa(row_num))
		}
		if row_num == 5 {
			// fmt.Printf("Break   \n")
			// Board is filled with 5 rows
			fmt.Printf("Board rows: %s \n", strconv.Itoa(len(board.rows)))
			game.boards = append(game.boards, board)
			board = Board{}
			row_num = 0
			continue
		}
	}
	return game
}

func checkWinCondition(board Board, cur_num string) bool {
	win := false
	// fmt.Printf("Drawn number = %s \n", cur_num)
	// Check horizontal first
	for i, r := range board.rows {
		if len(r.fields) == 0 {
			continue
		}
		// fmt.Println(r)
		// Loop over row and mark number and check win condition
		winning_row := true
		for ri, v := range r.fields {
			// Mark number if called
			// fmt.Printf(strconv.FormatBool(v.number == cur_num))
			if v.number == cur_num {
				fmt.Printf("Marked %s \n", r.fields[ri].number)
				r.fields[ri].marked = true
				v.marked = true
			}
			// If number remains unmarked set win to false for this row
			if r.fields[ri].marked == false {
				winning_row = false
			}
		}
		if winning_row == true {
			fmt.Printf("Winning row %s \n", strconv.Itoa(i+1))
			fmt.Println(r)
			win = true
		}

	}

	// Check vertical

	for i := 0; i < 5; i++ {
		// Go over the column
		winning_col := true
		for ir, r := range board.rows {
			if len(r.fields) == 0 {
				continue
			}
			// Check row ir col i
			if r.fields[i].number == cur_num {
				fmt.Printf("Marked %s in col %s row %s \n", r.fields[i].number, strconv.Itoa(i+1), strconv.Itoa(ir+1))
				r.fields[i].marked = true
			}
			if r.fields[i].marked == false {
				winning_col = false
			}
		}
		if winning_col == true {
			fmt.Printf("Winning col %s \n", strconv.Itoa(i+1))
			win = true
		}

	}
	// Again if any is false break loop
	return win

}

func getBingoNumbers(rawBytes []byte) []string {
	var numbers []string

	lines := strings.Split(string(rawBytes), "\n")
	for i, line := range lines {
		if i == 0 {
			numbers = strings.Split(line, ",")
		}
	}
	return numbers
}

func calculateScore(board Board) int {
	sum_unmarked := 0
	for _, r := range board.rows {
		if len(r.fields) == 0 {
			continue
		}
		// fmt.Println(r)
		// Loop over row and mark number and check win condition
		for _, v := range r.fields {
			// Mark number if called
			// fmt.Printf(strconv.FormatBool(v.number == cur_num))
			if v.marked == false {
				number, _ := strconv.Atoi(v.number)
				sum_unmarked += number
			}
		}
	}
	return sum_unmarked
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

	bingo_numbers := getBingoNumbers(rawBytes)
	// fmt.Println(bingo_numbers)

	game := createGame(rawBytes)
	fmt.Println(len(game.boards))
	game_won := false
	final_score := 0

	// Draw number and record it for calculating later
	for _, v := range bingo_numbers {
		fmt.Printf("Bingo number called: %s \n", v)

		// Check every board for number
		for i, b := range game.boards {
			// fmt.Println(b.rows)
			fmt.Printf("Checking board %s \n", strconv.Itoa(i+1))
			board_wins := checkWinCondition(b, v)
			if board_wins == true {
				fmt.Printf("\n Board %s wins \n", strconv.Itoa(i+1))
				fmt.Println(b)
				game_won = true
				score := calculateScore(b)
				fmt.Println(score)
				called_num, _ := strconv.Atoi(v)
				final_score = score * called_num
				break
			}
			if game_won == true {
				break
			}

		}
		if game_won == true {
			break
		}
	}
	fmt.Printf("Final score: %s \n", strconv.Itoa(final_score))

}
