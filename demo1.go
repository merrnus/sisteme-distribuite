package main

import "fmt"

func selam() {
	fmt.Println("salut de la Goroutine")
}

func main() {
	go selam()
	fmt.Println("a terminat main ul")
}
