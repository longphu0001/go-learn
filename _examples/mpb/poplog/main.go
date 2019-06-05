package main

import (
	"fmt"
	"io"
	"math/rand"
	"sync"
	"time"

	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	p := mpb.New(mpb.PopCompletedMode())

	total, numBars := 100, 2
	for i := 0; i < numBars; i++ {
		name := fmt.Sprintf("Bar#%d:", i)
		bar := p.AddBar(int64(total),
			mpb.BarNoPop(),
			mpb.PrependDecorators(
				decor.Name(name),
				decor.Percentage(decor.WCSyncSpace),
			),
			mpb.AppendDecorators(
				decor.OnComplete(
					decor.EwmaETA(decor.ET_STYLE_GO, 60), "done!",
				),
			),
		)
		// simulating some work
		go func() {
			max := 100 * time.Millisecond
			for i := 0; i < total; i++ {
				start := time.Now()
				time.Sleep(time.Duration(rand.Intn(10)+1) * max / 10)
				bar.Increment(time.Since(start))
			}
		}()
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		max := 3000 * time.Millisecond
		for i := 0; i < 10; i++ {
			p.Add(0, makeLogBar(fmt.Sprintf("some log: %d", i))).SetTotal(0, true)
			time.Sleep(time.Duration(rand.Intn(10)+1) * max / 10)
		}
	}()

	wg.Wait()
	p.Wait()
}

func makeLogBar(msg string) mpb.FillerFunc {
	limit := "%%.%ds"
	return func(w io.Writer, width int, st *decor.Statistics) {
		fmt.Fprintf(w, fmt.Sprintf(limit, width), msg)
	}
}
