package p1

import (
	"fmt"
	"time"
)

func p2() (result int) {
	defer timeTrack(time.Now(), "p2")
	list := parse()

	answer := 0

	for i, s := range list {
		find := 2020 - s
		for j, v := range list[i:] {
			if v < find {
				left_over := find - v
				for _, x := range list[i+j:] {
					if x == left_over {
						fmt.Println(s, v, x)
						answer = s * v * x
					}
				}
			}
		}

	}

	return answer

}

func main() {
	fmt.Println(p2())
}
