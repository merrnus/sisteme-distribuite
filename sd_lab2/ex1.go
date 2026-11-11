// ex1: Transfera un mesaj de la o goroutine la main printr-un channel.

package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "salut"
	}()

	fmt.Println(<-ch)
}

