package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

//Sleeper allows you to put delays
type Sleeper interface {
	Sleep()
}

//ConfigurableSleeper ConfigurableSleeper
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

//Sleep ...
func (c ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

//Countdown counts in opposite direction than up
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprintf(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
