package dump

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func goroutines() {
	f("direct")
	go f("goroutine")

	go func(msg string) {
		for i := 0; i < 3; i++ {
			fmt.Println(msg, ":", i)
		}
	}("goroutines")
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	time.Sleep(time.Second)
	fmt.Println("done")
}