package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const coundownStart = 3

type Sleeper interface {
  Sleep()
}

type ConfigurableSleeper struct {
  duration time.Duration
  sleep func(time.Duration)
}

type SpyTime struct {
  durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
  s.durationSlept = duration
}

type SpySleeper struct {
  Calls int
}

type SpyCountdownOperations struct {
  Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
  s.Calls = append(s.Calls, sleep)
}
func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
  s.Calls = append(s.Calls, write)
  return
}

func(c *ConfigurableSleeper) Sleep() {
  c.sleep(c.duration)
}

const write = "write"
const sleep = "sleep"

func (s *SpySleeper) Sleep() {
  s.Calls++
}


func Countdown(out io.Writer, sleeper Sleeper) {
  for i := coundownStart; i > 0; i-- {
    fmt.Fprintln(out, i)
    sleeper.Sleep()
  }
  fmt.Fprint(out, finalWord)
}

func main() {
  sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
  Countdown(os.Stdout, sleeper)
}
