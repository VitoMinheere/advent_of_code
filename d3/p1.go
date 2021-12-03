package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
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

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var gamma_list []string
	var epsilon_list []string
	var gamma_rate int64
	var epsilon_rate int64
	var pos int

	scanner := bufio.NewScanner(file)
	// file_length, err := lineCounter(file)
	// fmt.Println(file_length)
	bit_list := make([][]string, 1000, 1000)

	// Loop over depths
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, "")
		fmt.Println("Printing current Row", s)
		bit_list[pos] = append(bit_list[pos], s...)
		pos++
	}

	for j := 0; j < 12; j++ {
		var x []string
		for _, v := range bit_list {
			// Append the wanted char from slice
			x = append(x, v[j])
		}
		most_freq_bit := mostFrequent(x)
		least_freq_bit := "0"
		if most_freq_bit == "0" {
			least_freq_bit = "1"
		}
		gamma_list = append(gamma_list, most_freq_bit)
		epsilon_list = append(epsilon_list, least_freq_bit)

	}
	gamma := strings.Join(gamma_list, "")
	gamma_rate, err = strconv.ParseInt(gamma, 2, 64)
	fmt.Println("Gamma rate", gamma_rate)
	epsilon := strings.Join(epsilon_list, "")
	epsilon_rate, err = strconv.ParseInt(epsilon, 2, 64)
	fmt.Println("Epsilon rate", epsilon_rate)

	res := gamma_rate * epsilon_rate
	fmt.Println("Result ", res)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
