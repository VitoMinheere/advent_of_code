package p1

import (
	"fmt"
	"time"
)

func main() {
	defer timeTrack(time.Now(), "p1")
	list := parse()

	for i, s := range list {
		find := 2020 - s
		for _, v := range list[i:] {
			if v == find {
				fmt.Println(v * s)
				return
			}
		}
	}
}
