package main

import (
	"fmt"
	"time"
)

func main() {
	sw := Stopwatch{}
	sw.Start()

	time.Sleep(1 * time.Second)
	sw.SaveSplit()

	time.Sleep(500 * time.Millisecond)
	sw.SaveSplit()

	time.Sleep(300 * time.Millisecond)
	sw.SaveSplit()

	fmt.Println(sw.GetResults())
}

type Stopwatch struct {
	start  time.Time
	splits []time.Time
}

func (sw *Stopwatch) Start() {
	sw.start = time.Now()
	sw.splits = make([]time.Time, 0, 10)
}

func (sw *Stopwatch) SaveSplit() {
	sw.splits = append(sw.splits, time.Now())
}

func (sw *Stopwatch) GetResults() []time.Duration {
	durations := make([]time.Duration, len(sw.splits))
	for i, v := range sw.splits {
		durations[i] = v.Sub(sw.start)
	}
	return durations
}
