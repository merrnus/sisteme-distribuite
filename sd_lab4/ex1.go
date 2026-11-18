package main

import (
	"fmt"
	"sync"
)

type Msg struct {
	Value int
	Done  bool // liderul anunta ca a fost ales
}

// realizeaza operatiile interne nodului
func node(id int, in <-chan Msg, out chan<- Msg, isLeader *bool, wg *sync.WaitGroup) {
	defer wg.Done()
	out <- Msg{Value: id, Done: false}

	for msg := range in {
		if msg.Done {
			if !*isLeader {
				out <- msg
			}
			return
		}
		if msg.Value > id {
			out <- msg
		} else if msg.Value == id {
			*isLeader = true
			out <- Msg{Value: id, Done: true}
		}
	}
}

// returneaza id-ul liderului
func RingFlooding(ids []int) int {
	n := len(ids)
	chans := make([]chan Msg, n)
	for i := range chans {
		chans[i] = make(chan Msg, 2)
	}

	leaders := make([]bool, n)
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go node(ids[i], chans[i], chans[(i+1)%n], &leaders[i], &wg)
	}

	wg.Wait()

	for i, isL := range leaders {
		if isL {
			return ids[i]
		}
	}
	return -1
}

func main() {
	ids := []int{12, 45, 7, 89, 23}
	fmt.Println("Lider:", RingFlooding(ids))
}
