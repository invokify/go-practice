package pipeline

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

func TestRepeatFn(t *testing.T) {
	start := time.Now()

	done := make(chan int)
	defer close(done)

	//Stage 1: Generate random numbers
	randNumFetcher := func() int {
		return rand.Intn(500000000)
	}
	randIntStream := repeatFn(done, randNumFetcher)

	// Stage 2: Find prime numbers by fanning out
	CPUCount := runtime.NumCPU()
	primeFinderChannels := make([]<-chan int, CPUCount)
	for i := 0; i < CPUCount; i++ {
		primeFinderChannels[i] = primeFinder(done, randIntStream)
	}

	// Stage 3: Take 10 prime numbers by fanning in
	mergedStream := merge(done, primeFinderChannels...)
	for primeNum := range take(done, mergedStream, 10) {
		fmt.Println(primeNum)
	}

	elapsed := time.Since(start)
	fmt.Printf("Took %s\n", elapsed)
}
