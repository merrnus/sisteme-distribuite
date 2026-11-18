package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
)

type LamportClock struct {
	counter int
}

// incrementeaza contorul
func (l *LamportClock) Tick() {
	l.counter++
}

func (l *LamportClock) Send() int {
	l.counter++
	return l.counter
}

func (l *LamportClock) Receive(ts int) {
	if ts > l.counter {
		l.counter = ts
	}
	l.counter++
}

type Message struct {
	sender, ts int
}

type Event struct {
	internal_counter int
	process          int
	kind             string // "internal", "send", "recv"
}

// returneaza true daca a precede b
func less(a, b Event) bool {
	if a.internal_counter == b.internal_counter {
		return a.process < b.process
	}
	return a.internal_counter < b.internal_counter
}

func worker(id int, in <-chan Message, out chan<- Message, events chan<- Event, wg *sync.WaitGroup) {
	defer wg.Done()
	clock := LamportClock{counter: 0}

	for i := 0; i < rand.Intn(3)+1; i++ {
		clock.Tick()
		events <- Event{internal_counter: clock.counter, process: id, kind: "internal"}
	}

	if id == 1 {
		ts := clock.Send()
		events <- Event{internal_counter: ts, process: id, kind: "send"}
		out <- Message{sender: id, ts: ts}

		msg := <-in
		clock.Receive(msg.ts)
		events <- Event{internal_counter: clock.counter, process: id, kind: "recv"}
	} else {
		msg := <-in
		clock.Receive(msg.ts)
		events <- Event{internal_counter: clock.counter, process: id, kind: "recv"}

		for i := 0; i < rand.Intn(3)+1; i++ {
			clock.Tick()
			events <- Event{internal_counter: clock.counter, process: id, kind: "internal"}
		}

		ts := clock.Send()
		events <- Event{internal_counter: ts, process: id, kind: "send"}
		out <- Message{sender: id, ts: ts}
	}
}

func main() {
	c12 := make(chan Message, 1)
	c23 := make(chan Message, 1)
	c31 := make(chan Message, 1)
	eventsChan := make(chan Event, 30)

	var wg sync.WaitGroup
	wg.Add(3)

	go worker(1, c31, c12, eventsChan, &wg)
	go worker(2, c12, c23, eventsChan, &wg)
	go worker(3, c23, c31, eventsChan, &wg)

	wg.Wait()
	close(eventsChan)

	var events []Event
	for e := range eventsChan {
		events = append(events, e)
	}

	sort.Slice(events, func(i, j int) bool {
		return less(events[i], events[j])
	})

	for _, e := range events {
		fmt.Printf("P%d | ts: %2d | %s\n", e.process, e.internal_counter, e.kind)
	}
}
