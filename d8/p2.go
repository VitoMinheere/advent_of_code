package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CodeMapping struct {
	code  string
	value string
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// Decode what the extra wire is by checking which letter from y is not in x
func DecodeFromString(x string, y string) []string {
    var in_x bool
    var z []string

    split_y := strings.Split(y, "")
    for _, l := range split_y {
        in_x = strings.Contains(x, string(l))
        if in_x == false {
            z = append(z,l)
        }
    }
    return z
}


func MostFrequentString(arr []string) string {
    // Create string from list
    arr_string := strings.Join(arr, "")
    fmt.Printf("arr_string = %s \n", arr_string)
    cnt := 0
    var freq string

    for _, v := range arr {
        fmt.Printf("First char to check = %s \n", v)
        str_cnt := strings.Count(arr_string, v)
        if str_cnt > cnt {
            freq = v
        }
    }
    return freq
}

func removeDuplicateValues(stringSlice []string) []string {
    keys := make(map[string]bool)
    list := []string{}

    // If the key(values of the slice) is not equal
    // to the already present value in new slice (list)
    // then we append it. else we jump on another element.
    for _, entry := range stringSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }
    return list
}

func CheckIfAllContain(slice []string, x string) bool {
    res := true
    for _, v := range slice {
        if !strings.Contains(v, x) {
            res = false
        }
    }
    return res
}

func GetZero(c []string, eight string, middle_or_top_left []string) string {
    // var missing_wire []string
    zero := ""

    c = removeDuplicateValues(c)

    // loop over 0 6 9
    for _, v := range c {
        // Check that top_left is in all
        for _, char := range middle_or_top_left {
            if !strings.Contains(v, char) {
                zero = char
            }
        }
    }
    return zero
}

func MapSignal(c []string) [10]CodeMapping {
    var signal [10]CodeMapping
    var top string
    var middle_candidates []string
    var middle string
    var bottom string
    var top_right string
    var bottom_left string

    // Fill the CodeMapping once
    for _, v := range c {
		length := len(v)
        if length == 4 {
            signal[4] = CodeMapping{v, "4"}
        }
        if length == 3 {
            signal[7] = CodeMapping{v, "7"}
        }
        if length == 2 {
            signal[1] = CodeMapping{v, "1"}
        }
        if length == 7 {
            signal[8] = CodeMapping{v, "8"}
        }

    }
    two_and_five := []string{}
    //TODO Create a while loop to keep looping until all are set
    for _, v := range c {
        // fmt.Printf("%s \n", v)
		length := len(v)
        // Compare 1 and 7 to get top value
        if len(top) == 0 {
            top = DecodeFromString(signal[1].code, signal[7].code)[0]
            fmt.Printf("Top is %s \n", top)
        }

        // Could be a 0 6 or 9, get the middle by checking which value is not in 8
        if length == 5 {
            two_and_five = append(two_and_five, v)
        }
    }

    for _, v := range c {
        // fmt.Printf("%s \n", v)
		length := len(v)
        if length == 6 {
            if len(middle_candidates) < 3 {
                middle_candidates = append(middle_candidates, v)
            }
            if len(middle_candidates) == 3 {
                middle_candidates = append(middle_candidates, v)
                middle_or_top_left := DecodeFromString(signal[1].code, signal[4].code)
                middle = GetZero(middle_candidates, signal[8].code, middle_or_top_left)
                fmt.Printf("Middle is %s \n", middle)

                signal[0] = CodeMapping{strings.Replace(signal[8].code, middle, "", 1), "0"}

                // Check value with 4 to get top and bottom row back
                if len(bottom) == 0 {
                    bottom_slice := DecodeFromString(signal[4].code, signal[0].code)
                    // Either top, bottom or bottom left
                    if len(bottom_slice) > 0{
                        for _, bv := range bottom_slice {
                            if bv != top && len(bottom) == 0 {
                                // not top
                                fmt.Print("two and five")
                                fmt.Println(two_and_five)
                                if len(two_and_five) > 0 {
                                    contains_bv := CheckIfAllContain(two_and_five, bv)
                                    if contains_bv == true {
                                            fmt.Printf("%s contains %s \n", two_and_five, bv)
                                            bottom = bv
                                            fmt.Printf("Bottom is %s \n", bottom)
                                        }
                                    }
                                }
                            }
                        }
                    }
                }

                signal[9] = CodeMapping{signal[4].code+top+bottom, "9"}
            }

            }

    for _, v := range c {
		length := len(v)
        // Could be 3 or 5
        if length == 5 {
            // Determine if it is 5 by checking against 9
            missing_slice := DecodeFromString(v, signal[9].code)
            if len(missing_slice) == 1 {
                res := DecodeFromString(signal[1].code, missing_slice[0])
                fmt.Print("res")
                fmt.Println(res)
                // If res containts something the value was a 5 or a 2
                if len(res) > 0 {
                    bottom_left = DecodeFromString(signal[9].code, signal[8].code)[0]
                    fmt.Printf("Bottom left is %s \n", bottom_left)
                    // bottom_left = res[0]
                    // fmt.Printf("signal 6 = %s \n", signal[6].code)
                    signal[5] = CodeMapping{strings.Replace(signal[6].code, bottom_left, "", 1), "5"}
                } else {
                    fmt.Print("Missing slioce")
                    fmt.Println(missing_slice)
                    top_right = missing_slice[0]
                    fmt.Printf("Top right is %s \n", top_right)
                }
            }
        }
        signal[3] = CodeMapping{signal[7].code + bottom + middle, "3"}
        if len(top_right) == 1 {
            signal[6] = CodeMapping{strings.Replace(signal[8].code, top_right, "", 1), "6"}
            // bottom_right = strings.Replace(signal[1].code, top_right, "", 1)
        }

        bottom_left = DecodeFromString(signal[9].code, signal[8].code)[0]
        signal[2] = CodeMapping{top + middle + bottom + top_right + bottom_left, "2"}
    }

    for _, v := range c {
		length := len(v)
        // Could be 3 or 5
        if length == 5 {
            // Determine if it is 5 by checking against 9
            missing_slice := DecodeFromString(v, signal[9].code)
            if len(missing_slice) == 1 {
                res := DecodeFromString(signal[1].code, missing_slice[0])
                fmt.Print("res")
                fmt.Println(res)
                // If res containts something the value was a 5 or a 2
                if len(res) > 0 {
                    bottom_left = DecodeFromString(signal[9].code, signal[8].code)[0]
                    fmt.Printf("Bottom left is %s \n", bottom_left)
                    // bottom_left = res[0]
                    // fmt.Printf("signal 6 = %s \n", signal[6].code)
                    signal[5] = CodeMapping{strings.Replace(signal[6].code, bottom_left, "", 1), "5"}
                } else {
                    fmt.Print("Missing slioce")
                    fmt.Println(missing_slice)
                    top_right = missing_slice[0]
                    fmt.Printf("Top right is %s \n", top_right)
                }
            }
        }
    }
    return signal
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// First scan outside of loop to get first depth
	simple_digits := 0
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, "| ")
		code := s[0]
        input_code := strings.Split(code, " ")
        o_code := s[1]
        output_code := strings.Split(o_code, " ")
		fmt.Println(output_code)
        codes := MapSignal(input_code)
        fmt.Print(codes)
		var output_slice []string
		for _, v := range output_code {
			value := ""
			length := len(v)
			if length == 2 {
				value = "1"
			}
			if length == 3 {
				value = "7"
			}
			if length == 4 {
				value = "4"
			}
			if length == 7 {
				value = "8"
			}
			// fmt.Printf("Length of string %s \n", strconv.Itoa(length))
			if value == "" {
				sorted_v := SortString(v)
				// fmt.Printf("sorted_v: %s \n", sorted_v)
				for _, c := range codes {
					sorted_c := SortString(c.code)
					// fmt.Printf("sorted_c: %s \n", sorted_c)
					if sorted_v == sorted_c {
						// fmt.Printf("Code found: %s \n", c.code)
						// fmt.Printf("Value found: %s \n", c.value)
						value = c.value
					}
				}

			}

			// fmt.Printf("Value = %s \n", value)
			output_slice = append(output_slice, value)
		}
		output := strings.Join(output_slice, "")
		fmt.Printf("\n Output = %s \n", output)
		output_val, _ := strconv.Atoi(output)
		simple_digits += output_val
	}
	fmt.Printf("Simple digits: %s \n", strconv.Itoa(simple_digits))

}
