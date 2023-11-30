package main

import "fmt"

func main() {
	defer timeTrack(time.Now(), "p1")
	list := parse()
	fmt.Print(list)

}
