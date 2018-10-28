package bug3

import (
	"fmt"
	"time"
)

func producer2(ch chan uint) {
	for i := uint(0); i < 2; i++ {
		ch <- i // Send i to ch
	}
	close(ch)
}

func producer5(ch chan uint) {
	for i := uint(0); i < 5; i++ {
		ch <- i // Send i to ch
	}
	close(ch)
}

func printAll(ch chan uint) {
	for x := range ch {
		fmt.Println(x)
	}
}

func Example2Bug3() {
	out := bug3(producer2, producer2)
	go printAll(out)
	<-time.NewTimer(1 * time.Millisecond).C
	// Unordered Output: 0
	// 1
	// 0
	// 1
}

func Example5Bug3() {
	out := bug3(producer5, producer5)
	go printAll(out)
	<-time.NewTimer(1 * time.Millisecond).C
	// Unordered Output: 0
	// 1
	// 2
	// 3
	// 4
	// 0
	// 1
	// 2
	// 3
	// 4
}

func Example25Bug3() {
	out := bug3(producer2, producer5)
	go printAll(out)
	<-time.NewTimer(1 * time.Millisecond).C
	// Unordered Output: 0
	// 1
	// 0
	// 1
	// 2
	// 3
	// 4
}
