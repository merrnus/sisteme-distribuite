// ex3: main primeste doua mesaje pe acelasi canal.

package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "mesajul 1"
		ch <- "mesajul 2"
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
