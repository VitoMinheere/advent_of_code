package main

import (
    "os"
    "bufio"
    "bytes"
    "io"
    "fmt"
    "strconv"
)

// Read a whole file into the memory and store it as array of lines
func readLines(path string) (lines []string, err error) {
    var (
        file *os.File
        part []byte
        prefix bool
    )
    if file, err = os.Open(path); err != nil {
        return
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    buffer := bytes.NewBuffer(make([]byte, 0))
    for {
        if part, prefix, err = reader.ReadLine(); err != nil {
            break
        }
        buffer.Write(part)
        if !prefix {
            lines = append(lines, buffer.String())
            buffer.Reset()
        }
    }
    if err == io.EOF {
        err = nil
    }
    return
}

func findArraySum(arr []string) int{
   res := 0
   for i:=0; i<len(arr); i++ {
       str_int, _ := strconv.Atoi(arr[i])
       res += str_int
   }
   return res
}

func main() {
    lines, err := readLines("input.txt")
    if err != nil {
        fmt.Println("Error: %s\n", err)
        return
    }
    var count int
    for i, _ := range lines {
        if i+3 < len(lines){

            cur := lines[i:i+3]
            fmt.Println(cur)
            next := lines[i+1:i+4]
            fmt.Println(next)

            sum_cur := findArraySum(cur)
            sum_next := findArraySum(next)
            fmt.Println(sum_cur)
            fmt.Println(sum_next)
            if (sum_cur - sum_next) < 0{
                count++
            }

        }
    }
    fmt.Println(count)
}
