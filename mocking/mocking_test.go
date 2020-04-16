package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("time and data", func(t *testing.T) {

		buffer := &bytes.Buffer{}
		Countdown(buffer, &CountdownOperationsSpy{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("Wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSleppt != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime)
	}
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

type SpyTime struct {
	durationSleppt time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSleppt = duration
}
