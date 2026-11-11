// ex4: O goroutine notifica main cand termina, fara WaitGroup.

package main

import "fmt"

func main() {
	done := make(chan bool)

	go func() {
		fmt.Println("proccessing...")
		done <- true
	}()

	<-done
	fmt.Println("a terminat goroutine, main ul continue")
}
