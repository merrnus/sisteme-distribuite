package main

import (
	"fmt"
	"math/rand"
	"time"
)

func robot(name string, finish chan<- string) {
	for step := 1; step <= 10; step++ {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
	finish <- name
}

func main() {
	r1 := make(chan string, 1)
	r2 := make(chan string, 1)

	go robot("Robot 1", r1)
	go robot("Robot 2", r2)

	select {
	case w := <-r1:
		fmt.Println("Primul la finish:", w)
	case w := <-r2:
		fmt.Println("Primul la finish:", w)
	}
}
