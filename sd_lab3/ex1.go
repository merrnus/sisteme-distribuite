package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// stocheaza deviatia fata de timpul real
type Clock struct {
	offset time.Duration
}

func (c Clock) Now() time.Time {
	return time.Now().Add(c.offset)
}

// Calculeaza decalajul theta si round-trip-ul delta
func estimate(t1, t2, t3, t4 time.Time) (theta, delta time.Duration) {
	delta = t4.Sub(t1) - t3.Sub(t2)
	theta = (t2.Sub(t1) + t3.Sub(t4)) / 2
	return theta, delta
}

func measure(c_clock, s_clock Clock) (theta, delta time.Duration) {
	// Request
	t1 := c_clock.Now()
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t2 := s_clock.Now()

	// Procesare intre 5 si 20 ms
	time.Sleep(5*time.Millisecond + time.Duration(rand.Intn(20))*time.Millisecond)

	t3 := s_clock.Now()
	// Response
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t4 := c_clock.Now()

	return estimate(t1, t2, t3, t4)
}

func main() {
	s_clock := Clock{offset: 0}
	c_clock := Clock{offset: 500 * time.Millisecond}

	// (a)
	thetaA, deltaA := measure(c_clock, s_clock)
	fmt.Printf("a) theta: %v, delta: %v, offset nou: %v\n", thetaA, deltaA, c_clock.offset+thetaA)

	// (b)
	minDelta := time.Duration(math.MaxInt64)
	bestTheta := time.Duration(0)

	for i := 0; i < 10; i++ {
		theta, delta := measure(c_clock, s_clock)
		if delta < minDelta {
			minDelta = delta
			bestTheta = theta
		}
	}

	fmt.Printf("b) min delta: %v, best theta: %v, offset nou: %v\n", minDelta, bestTheta, c_clock.offset+bestTheta)
}
