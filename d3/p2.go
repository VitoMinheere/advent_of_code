package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
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
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func mostFrequent(arr []string) string { // assuming no tie
	var cnt1 int
	var cnt0 int

	for _, a := range arr {
		if a == "0" {
			cnt0++
		} else {
			cnt1++
		}
	}
	freq := "0"
	if cnt1 > cnt0 {
		freq = "1"
	}

	return freq
}

func getFrequentFromSlice(bit_list [][]string) (string, string) {
	var most_freq_bit string
	var least_freq_bit string

	for j := 0; j < 12; j++ {
		var x []string
		for _, v := range bit_list {
			// Append the wanted char from slice
			x = append(x, v[j])
		}
		most_freq_bit = mostFrequent(x)
		least_freq_bit = "0"
		if most_freq_bit == "0" {
			least_freq_bit = "1"
		}

	}
	return most_freq_bit, least_freq_bit
}

func createMatrix(file *os.File) [][]string {
	var pos int
	scanner := bufio.NewScanner(file)
	file_length, _ := lineCounter(file)
	bit_list := make([][]string, file_length, file_length)

	// Loop over depths
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, "")
		bit_list[pos] = append(bit_list[pos], s...)
		pos++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return bit_list

}

func remove(s [][]string, i int) [][]string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// var gamma_list []string
	// var epsilon_list []string
	// var gamma_rate int64
	// var epsilon_rate int64
	var pos int

	// b_list := createMatrix(file)
	// fmt.Println(b_list)
	scanner := bufio.NewScanner(file)
	// file_length, err := lineCounter(file)
	bit_list := make([][]string, 1000, 1000)

	// Loop over depths
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, "")
		bit_list[pos] = append(bit_list[pos], s...)
		pos++
	}

	bit_loc := 0
	smaller_list := bit_list
	for len(smaller_list) >= 1 && bit_loc < 12 {
		most, _ := getFrequentFromSlice(bit_list)
		for i, v := range smaller_list {
			if v[bit_loc] != most {
				smaller_list = remove(bit_list, i)
			}

		}
		fmt.Println(len(smaller_list))
		bit_loc++
	}
	fmt.Println("Oxygen list", len(smaller_list))

	// for j := 0; j < 12; j++ {
	// 	var x []string
	// 	for _, v := range bit_list {
	// 		// Append the wanted char from slice
	// 		x = append(x, v[j])
	// 	}
	// 	most_freq_bit := mostFrequent(x)
	// 	least_freq_bit := "0"
	// 	if most_freq_bit == "0" {
	// 		least_freq_bit = "1"
	// 	}
	// 	gamma_list = append(gamma_list, most_freq_bit)
	// 	epsilon_list = append(epsilon_list, least_freq_bit)

	// }
	// gamma := strings.Join(gamma_list, "")
	// gamma_rate, err = strconv.ParseInt(gamma, 2, 64)
	// fmt.Println("Gamma rate", gamma_rate)
	// epsilon := strings.Join(epsilon_list, "")
	// epsilon_rate, err = strconv.ParseInt(epsilon, 2, 64)
	// fmt.Println("Epsilon rate", epsilon_rate)

	// res := gamma_rate * epsilon_rate
	// fmt.Println("Result ", res)

}
