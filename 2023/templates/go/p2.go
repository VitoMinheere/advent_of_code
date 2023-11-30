package main

import "fmt"

func main() {
	defer timeTrack(time.Now(), "p2")
	list := parse()
	fmt.Print(list)

}

