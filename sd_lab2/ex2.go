// ex2: Transfera un mesaj de la main la o goroutine printr-un channel.

package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan string)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(<-ch)
	}()

	ch <- "salut"
	wg.Wait()
}
