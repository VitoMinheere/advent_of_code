package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var cur_depth int
    var next_depth int
    var num_increases int

    scanner := bufio.NewScanner(file)
    // First scan outside of loop to get first depth
    scanner.Scan()
    cur_depth, err = strconv.Atoi(scanner.Text())
    fmt.Println(cur_depth)

    // Loop over depths
    for scanner.Scan() {
        next_depth, err = strconv.Atoi(scanner.Text())
        fmt.Println(next_depth)
        if (cur_depth - next_depth) < 0{
            num_increases++
        }
        cur_depth = next_depth
        fmt.Println(num_increases)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
